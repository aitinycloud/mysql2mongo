package handle

import (
	"database/sql"
	"fmt"
	"os"

	"../pkg/logging"
	_ "github.com/siddontang/go-mysql/driver"
	//_ "github.com/go-sql-driver/mysql"
)

var BinlogName string
var BinlogPos uint32

func GetOnlineMasterBinInfo(dbname string, username string, passwd string, host string) (string, uint32) {
	var binname string
	var binpos uint32
	defer func() {
		if err := recover(); err != nil {
			logging.Error("mysql get master binlog infomation error , panic : ", err)
			os.Exit(-1)
		}
	}()
	sqlconnstr := fmt.Sprintf("%s:%s@%s?%s", username, passwd, host, dbname)
	db, err := sql.Open("mysql", sqlconnstr)
	checkErr(err)
	//查询数据
	sqlstr := `SHOW MASTER STATUS`
	rows, err := db.Query(sqlstr)
	checkErr(err)
	for rows.Next() {
		var strFile string
		var uintPosition uint32
		var strBinlog_Do_DB string
		var strBinlog_Ignore_DB string
		rows.Scan(&strFile, &uintPosition, &strBinlog_Do_DB, &strBinlog_Ignore_DB)
		binname = strFile
		binpos = uintPosition
		SetMasterCacheBinlogPos(binname, binpos)
	}
	rows.Close()
	db.Close()
	return binname, binpos
}

func GetMasterCacheBinlogPos() (string, uint32) {
	return BinlogName, BinlogPos
}

func SetMasterCacheBinlogPos(binlogname string, binlogpos uint32) {
	BinlogName = binlogname
	BinlogPos = binlogpos
	//setting.BinlogName = binlogname
	//setting.BinlogPos = binlogpos
}

func SetMasterBinlogPosToConfig(binlogname string, binlogpos uint32) {
	SetMasterCacheBinlogPos(binlogname, binlogpos)
	//setting.SetBinlog(binlogname, binlogpos)
}
