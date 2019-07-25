package handle

import (
	"database/sql"
	"time"

	"../pkg/logging"

	_ "github.com/siddontang/go-mysql/driver"
)

//Master InsightDB
var MasterInsightDB *sql.DB


func Setup() {
	logging.Info("Handle Setup.")
	SyncInit()
}

func GetServerConfig() {
	
}

func ClearData() {
	
	//logging.Info("ClearData END.")
}

func FullSyncDBTable(dbname string, tablename string, sqlstr string) {
	
}

func InreaseSyncHandle() {
	//1 start sync mysql binlog.
	logging.Info("start sync mysql binlog.")
}

//DoSync do work Sync.
func DoSync() {
	time.Sleep(100 * time.Millisecond)
	//Is binlog post setting.
	var NeedFullSync bool = true
	logging.Info(setting.BinlogPos, " , ", setting.BinlogName)
	if setting.BinlogPos != 0 && setting.BinlogName != "" {
		NeedFullSync = false
	}
	if NeedFullSync {
		logging.Info("NeedFullSync. Start Full Sync.")
		ClearData()
		//

		logging.Info("Full sync tables success.")
	}
	//increase sync.
	go InreaseSyncHandle()
}

//Work handle work.
func Work() {
	go DoSync()
}
