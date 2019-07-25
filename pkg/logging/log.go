package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

//Level log Level
type Level int

var (
	//F log f file
	F *os.File
	//DefaultPrefix Default Prefix
	DefaultPrefix = ""
	//DefaultCallerDepth Default CallerDepth
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

var (
	//ProjectName Project Name
	ProjectName = "insight"
	//ServerName Server Name
	ServerName = "mysqlsync"
	//ModuleName Module Name
	ModuleName = "handle"
)

const (
	//DEBUG log DEBUG
	DEBUG Level = iota
	//INFO log INFO
	INFO
	//WARNING log WARNING
	WARNING
	//ERROR log ERROR
	ERROR
	//FATAL log FATAL
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, 0)
}

//Debug log Debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

//Info log Debug level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

//Warn log Debug level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

//Error log Debug level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

//Fatal log Debug level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s][%s:%d][%s][%s][%s]", levelFlags[level], time.Now().Format(DateTimeFormat),
			filepath.Base(file), line, ProjectName, ServerName, ModuleName)
	} else {
		logPrefix = fmt.Sprintf("[%s][%s][:][%s][%s][%s]", levelFlags[level], time.Now().Format(DateTimeFormat),
			ProjectName, ServerName, ModuleName)
	}
	logger.SetPrefix(logPrefix)
}
