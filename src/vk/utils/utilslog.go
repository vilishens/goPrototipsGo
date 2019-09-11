package utils

import (
	"log"
	"os"
	vomni "vk/omnibus"
)

// String into a logger -- <PREFIX> <SEPARATOR> <DATE+TIME> <SEPARATOR> <STR>
func LogStr(d *log.Logger, str string) {
	if nil != d {
		strNew := vomni.UDPMessageSeparator + " " + str
		d.Println(strNew)
	}
}

// Logger with a prefix -- <PREFIX> <SEPARATOR> <DATE+TIME>
func LogNew(d *os.File, prefix string) (newLog *log.Logger) {
	return log.New(d, prefix+" "+vomni.UDPMessageSeparator+" ", vomni.LogLoggerFlags)
}

/*
// Point logger with no prefix -- <SEPARATOR> <DATE+TIME>
func LogNewPoint(d *os.File) (newLog *log.Logger) {
	return log.New(d, vomni.UDPMessageSeparator+" ", vomni.LogLoggerFlags)
}
*/

func LogNewPath(path string, prefix string) (newLog *log.Logger, err error) {

	f, err := OpenFile(path, vomni.LogFileFlags, vomni.LogUserPerms)
	if nil == err {
		return LogNew(f, prefix), err
	}

	return nil, err
}

func LogErr(err error) {

	LogStr(vomni.LogErr, err.Error())

}

func LogInfo(str string) {

	LogStr(vomni.LogInfo, str)

}

func LogData(str string) {

	LogStr(vomni.LogData, str)

}

func LogFatal(str string) {

	LogStr(vomni.LogFatal, str)

}
