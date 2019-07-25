package setting

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"github.com/gookit/config"
	"github.com/gookit/config/json"
)

const CONFIGFILENAME = "conf/conf.json"

var (
	//SelfIP Get Self IP
	SelfIP string

	//master config
	MasterType string
	//MasterDBName mysql DB Name
	MasterDBName string
	//MasterUser mysql user name
	MasterUser string
	//MasterPassword mysql password
	MasterPassword string
	//MasterHost mysql host IP:port
	MasterHost string

	//mongo config
	//SlaveDBName mysql DB Name
	MongoDBName string
	//SlaveUser mysql user name
	MongoUser string
	//SlavePassword mysql password
	MongoPassword string
	//SlaveHost mysql host IP:port
	MongoHost string

	//Binlog
	BinlogName string
	BinlogPos  uint32
)

//Setup setting Setup for load config.
func Setup() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(json.Driver)
	err := config.LoadFiles(CONFIGFILENAME)
	if err != nil {
		panic(err)
	}

	//binlog
	BinlogName = config.String("masterbinloginfo.bin_name")
	BinlogPos = 0

	//BinlogPos config DEBUG
	//buf := new(bytes.Buffer)
	//config.DumpTo(buf, config.JSON)
	//js, err := simplejson.NewJson([]byte(buf.String()))
	//if err != nil {
	//	panic(err)
	//}
	//tmpbin_pos, _ := js.Get("masterbinloginfo").Get("bin_pos").Int64()
	//BinlogPos = uint32(tmpbin_pos)

	MasterType = config.String("mastermysql.type")
	MasterDBName = config.String("mastermysql.dbname")
	MasterUser = config.String("mastermysql.user")
	MasterPassword = config.String("mastermysql.password")
	MasterHost = config.String("mastermysql.host")
}

func SetBinlog(name string, pos uint32) {
	BinlogName = name
	BinlogPos = pos
	config.Set("masterbinloginfo.bin_name", name)
	config.Set("masterbinloginfo.bin_pos", pos)
	buf := new(bytes.Buffer)
	config.DumpTo(buf, config.JSON)
	err := ioutil.WriteFile(CONFIGFILENAME, []byte(buf.String()), 0644) // oct, not hex
	if err != nil {
		fmt.Printf("File Write Error: %s\n", err)
		return
	}
}
