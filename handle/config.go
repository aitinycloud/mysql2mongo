package handle

import (
	"../pkg/logging"
	"../pkg/pubsub"
	"github.com/siddontang/go-mysql/mysql"
)

type DBServerInfo struct {
	DBtype string
	User   string
	Passwd string
	Host   string
	DBname string
}

type GroupIdToRegion struct {
	GroupId  string
	UserName string
	Region   string
}

type SqlInfoT struct {
	//for tye
	Retry     int
	DBname    string
	Tablename string
	Handle    string
	Content   map[string]string
}

type posSaver struct {
	pos   mysql.Position
	force bool
}

const (
	INSERT = "insert"
	UPDATE = "update"
	DELETE = "delete"
)

// pub sub topic.
const TOPIC = "synctopic"

//MSGPB
const PUBSUBMAXCOUNT = 6400

var MSGPB *pubsub.PubSub

const SERVERID = 100
const SqlConnMax = 16

func init() {
	MSGPB = pubsub.New(PUBSUBMAXCOUNT)
}

func checkErr(err error) {
	if err != nil {
		logging.Error(err)
	}
}

func checkErrPanic(err error) {
	if err != nil {
		logging.Error(err)
		panic(err)
	}
}
