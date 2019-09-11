package omnibus

import (
	"log"
	"os"
)

var stepList map[string]bool

var RootErr = make(chan error)
var RootDone = make(chan int)
var StepErr = make(chan error)

var (
	RootPath    string
	LogMainFile *os.File
	LogData     *log.Logger
	LogErr      *log.Logger
	LogFatal    *log.Logger
	LogInfo     *log.Logger
)

var PointLogData map[int]LogPointData
var PointCfgData map[int]CfgPointData

var MessageNumber int // unique message number (starting from the application launch)

var AllMessages map[int]MessageData

var CfgListSequence []int

func init() {

	PointCfgData = make(map[int]CfgPointData)
	PointCfgData[CfgTypeRelayInterval] = CfgPointData{CfgCd: CfgTypeRelayInterval, CfgStr: "relay-interval"}

	PointLogData = make(map[int]LogPointData)
	PointLogData[LogFileCdData] = LogPointData{LogCd: LogFileCdData, FileEnd: LogFileEndData, LogPrefix: LogPointPrefixData}
	PointLogData[LogFileCdErr] = LogPointData{LogCd: LogFileCdErr, FileEnd: LogFileEndErr, LogPrefix: LogPointPrefixErr}
	PointLogData[LogFileCdInfo] = LogPointData{LogCd: LogFileCdInfo, FileEnd: LogFileEndInfo, LogPrefix: LogPointPrefixInfo}

	CfgListSequence = []int{CfgTypeRelayInterval}
}
