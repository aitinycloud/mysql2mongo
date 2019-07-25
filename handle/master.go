// Project :  mysqlsync
// package :  handle
// file    :  master.go
// Get Mysql master logbin information and save it.
// Copyright (c) 2018-2018
// vixtel.com All rights reserved.

package handle

import (
	"sync"
	"time"
	"github.com/siddontang/go-mysql/mysql"
)

type masterInfo struct {
	sync.RWMutex
	Name         string
	Pos          uint32
	filePath     string
	fileContent  []byte
	lastSaveTime time.Time
}

func InitMasterInfo() (*masterInfo, error) {
	configfile := "cfg/conf.json"
	var m masterInfo
	bin_name, bin_pos := setting.BinlogName, setting.BinlogPos

	m.filePath = configfile
	m.lastSaveTime = time.Now()
	if bin_name != "" && bin_pos != 0 {
		m.Name = bin_name
		m.Pos = uint32(bin_pos)
	}
	return &m, nil
}

func (m *masterInfo) Save(pos mysql.Position) error {
	time.Sleep(time.Second)
	m.Lock()
	defer m.Unlock()

	m.Name = pos.Name
	m.Pos = pos.Pos
	if len(m.filePath) == 0 {
		return nil
	}
	n := time.Now()
	if n.Sub(m.lastSaveTime) < time.Second {
		return nil
	}
	m.lastSaveTime = n

	//setting.SetBinlog(m.Name, m.Pos)
	SetMasterBinlogPosToConfig(m.Name, m.Pos)
	return nil
}

func (m *masterInfo) Position() mysql.Position {
	m.RLock()
	defer m.RUnlock()
	return mysql.Position{
		Name: m.Name,
		Pos:  m.Pos,
	}
}

func (m *masterInfo) Close() error {
	pos := m.Position()
	return m.Save(pos)
}

func (m *masterInfo) GetConfigMasterBinlogPos() (string, uint32) {
	return m.Name, m.Pos
}
