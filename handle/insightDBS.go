// Project :  mysqlsync
// package :  handle
// file    :  insightDBS.go
// DB for canal handle.
// Copyright (c) 2018-2018
// vixtel.com All rights reserved.

package handle

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"../pkg/logging"

	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
)

type handler struct {
	canal.DummyEventHandler
	master *masterInfo
	syncCh chan interface{}
}

func (h *handler) OnDDL(nextPos mysql.Position, queryEvent *replication.QueryEvent) error {
	h.syncCh <- posSaver{nextPos, true}
	return nil
}

func (h *handler) OnXID(pos mysql.Position) error {
	h.syncCh <- posSaver{pos, false}
	return nil
}

func (h *handler) OnGTID(GTIDSet mysql.GTIDSet) error {
	return nil
}

func (h *handler) OnPosSynced(pos mysql.Position, tflag bool) error {
	return nil
}

func (h *handler) OnRotate(e *replication.RotateEvent) error {
	pos := mysql.Position{
		Name: string(e.NextLogName),
		Pos:  uint32(e.Position),
	}
	h.syncCh <- posSaver{pos, true}
	return nil
}

func (h *handler) OnTableChanged(schema string, table string) error {
	//fmt.Printf("OnTableChanged : \n")
	return nil
}

var GCOUTMutex sync.RWMutex

func (h *handler) OnRow(e *canal.RowsEvent) error {
	//log.Println("DB : ", e.Table.Schema, " table : ", e.Table.Name, " handle : ", e.Action)
	if !isSyncDBandTables(e.Table.Schema, e.Table.Name) {
		return nil
	}
	GCOUTMutex.Lock()
	defer GCOUTMutex.Unlock()
	arrcount := len(e.Rows)
	if e.Action == canal.InsertAction {
		for k := 0; k < arrcount; k++ {
			sqlinfo := SqlInfoT{}
			sqlinfo.DBname = e.Table.Schema
			sqlinfo.Tablename = e.Table.Name
			sqlinfo.Handle = ""
			sqlinfo.Content = make(map[string]string)
			for i := 0; i < len(e.Table.Columns); i++ {
				strvalue := fmt.Sprintf("%v", e.Rows[k][i])
				sqlinfo.Content[e.Table.Columns[i].Name] = strvalue
			}
			switch e.Action {
			case canal.InsertAction:
				sqlinfo.Handle = INSERT
			case canal.DeleteAction:
				sqlinfo.Handle = DELETE
			case canal.UpdateAction:
				sqlinfo.Handle = UPDATE
			default:
			}
		}
	}
	return nil
}

func (h *handler) String() string {
	return "TestHandler"
}

func (h *handler) syncLoop() {
	ticker := time.NewTicker(3000 * time.Millisecond)
	defer ticker.Stop()
	lastSavedTime := time.Now()
	var pos mysql.Position
	for {
		needFlush := false
		needSavePos := false
		select {
		case v := <-h.syncCh:
			switch v := v.(type) {
			case posSaver:
				now := time.Now()
				if v.force || now.Sub(lastSavedTime) > 3*time.Second {
					lastSavedTime = now
					needFlush = true
					needSavePos = true
					pos = v.pos
				}
			}
		case <-ticker.C:
			needFlush = true
		}
		if needFlush {

		}
		if needSavePos {
			if err := h.master.Save(pos); err != nil {
				logging.Info(fmt.Sprintf("save sync position %s err %v, close sync", pos, err))
				return
			}
		}
	}
}

func InsightDBInit(host string, user string, pwd string) {
	logging.Info("Init DB ", host, " ", user)
	var dbs = ""
	var tables = ""
	var tableDB = ""
	var ignoreTables = ""

	cfg := canal.NewDefaultConfig()
	cfg.Addr = fmt.Sprintf("%s", host)
	cfg.User = user
	cfg.Password = pwd
	cfg.Flavor = "mysql"

	cfg.ReadTimeout = 0 * 60 * 60 * time.Second
	cfg.HeartbeatPeriod = 0 * 60 * time.Second
	cfg.ServerID = uint32(SERVERID + 1)
	cfg.Dump.ExecutionPath = ""
	cfg.Dump.DiscardErr = false

	c, err := canal.NewCanal(cfg)
	if err != nil {
		fmt.Printf("create canal err %v", err)
		os.Exit(1)
	}

	if len(ignoreTables) == 0 {
		subs := strings.Split(ignoreTables, ",")
		for _, sub := range subs {
			if seps := strings.Split(sub, "."); len(seps) == 2 {
				c.AddDumpIgnoreTables(seps[0], seps[1])
			}
		}
	}

	if len(tables) > 0 && len(tableDB) > 0 {
		subs := strings.Split(tables, ",")
		c.AddDumpTables(tableDB, subs...)
	} else if len(dbs) > 0 {
		subs := strings.Split(dbs, ",")
		c.AddDumpDatabases(subs...)
	}
	mHandle := &handler{}
	mHandle.master, _ = InitMasterInfo()
	mHandle.syncCh = make(chan interface{}, 1)
	c.SetEventHandler(mHandle)
	go mHandle.syncLoop()
	// save Pos
	startName, startBinPos := mHandle.master.GetConfigMasterBinlogPos()
	startPos := mysql.Position{
		Name: startName,
		Pos:  uint32(startBinPos),
	}
	logging.Info("startName : ", startName, " startBinPos : ", startBinPos)
	go func() {
		err = c.RunFrom(startPos)
		if err != nil {
			fmt.Printf("start canal err %v", err)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sc
	c.Close()
	//LOG.INFO("Cannal InsightDB exit.")
}

func isSyncDBandTables(db string, table string) bool {

	return true
}
