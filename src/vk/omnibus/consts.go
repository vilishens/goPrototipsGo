package omnibus

import (
	"log"
	"os"
	"time"
)

const (
	PointCmdBits             = 0xFFF0000
	PointCmdOptionBits       = 0x000FFFF
	PointCmdLoadCfgIntoPoint = 0x0010000
	PointCmdSaveCfg          = 0x0020000
	PointCmdStopCfg          = 0x0040000
	PointCmdExitCfg          = 0x0080000
	PointCmdFreezeOn         = 0x0100000
	PointCmdFreezeOff        = 0x0200000
)

// constants for log
const (
	LogFileFlags       = os.O_RDWR | os.O_CREATE | os.O_APPEND
	LogUserPerms       = os.FileMode(0666)
	LogMainPath        = "../log/main/logMain.log"
	LogLoggerFlags     = log.LstdFlags | log.LUTC
	LogPrefixData      = "==== DATA ==="
	LogPrefixErr       = "!!! ERROR !!!"
	LogPrefixInfo      = "**** INFO ***"
	LogPrefixFatal     = "xxx FATAL xxx"
	LogFileCdErr       = 0x00004
	LogFileCdInfo      = 0x00001
	LogFileCdData      = 0x00002
	LogFileEndErr      = "err"
	LogFileEndInfo     = "info"
	LogFileEndData     = "data"
	LogPointPrefixData = "data"
	LogPointPrefixErr  = "err"
	LogPointPrefixInfo = "info"
)

//type LogCfg struct {
//	File string
//	List []string
//}

const (
	DoneError        = 0x0000010
	DoneReboot       = 0x0000020
	DoneRestart      = 0x0000040
	DoneUpdateGo     = 0x0000041
	DoneStop         = 0x0000080
	DonePostStop     = 0x0000100
	DoneDisconnected = 0x0000200
	DoneShutdown     = 0x0000400
	DoneExit         = 0x0000800
)

const (
	NoNetError           = 0x0000
	NoNetInternal        = 0x0010
	NoNetExternal        = 0x0020
	NetExternalNone      = 0x0000
	NetExternalNice2Have = 0x0001
	NetExternalRequired  = 0x0002
	NetExternalBits      = 0x0003
)

const (
	DelayStepExec             = 10 * time.Millisecond
	DelaySendMessage          = time.Millisecond // time delay between two message send
	DelaySendMessageListEmpty = 3 * time.Millisecond
	DelaySendMessageRepeat    = 500 * time.Millisecond // interval between repeated messages

	DelayWaitMessage = time.Millisecond // time delay between two message waiting

	DelayBetweenIPHello = 1000 * time.Millisecond

	MessageSendRepeatLimit = 3
)

const (
	StepNameConfig      = "step-config"
	StepNameMessages    = "step-messages"
	StepNameNetInfo     = "step-net-info"
	StepNameNetScan     = "step-net-scan"
	StepNameParams      = "step-params"
	StepNamePointConfig = "step-point-cfg"
	StepNamePointReady  = "step-point-ready"
	StepNamePointRun    = "step-point-run"
	StepNameRotateMain  = "step-rotate-main"
	StepNameStart       = "step-start"
	StepNameUDP         = "step-udp"
	StepNameWeb         = "step-web"
)

const (
	DirPermissions         = 0744
	FileNonExecPermissions = 0666
	FilePermissions        = 0644
)

const (
	TimeFormat1 = "2006-01-02 15:04:05 -07:00 MST"
)

const (
	CfgDefaultPath      = "../cfg/app/default.cfg"
	CliCfgPathFld       = "path"
	LogRotateStatusFile = "logStatus.status"
)

//############################################################################################
//################################################# Point run state codes ####################
//############################################################################################
const (
	PointStateUnknown      = 0x000000
	PointStateActive       = 0x000001
	PointStateSigned       = 0x000002
	PointStateDisconnected = 0x000004
	PointStateStoppingNow  = 0x000008

	PointCfgStateUnknown     = 0x000000
	PointCfgStateReady       = 0x000001
	PointCfgStateUnavailable = 0x000002

	PointNonActiveIndex = -1
)

const ()

//################################################# Configuration parameters ####################
const (
	CfgTypeUnknown       = 0x000000
	CfgTypeRelayInterval = 0x000001
	CfgTypeTempRelay     = 0x000002
)

//################################################# Net Info parameters (net/netinfo.go) ####################
const (
	NetInfoRepeats  = 3
	NetInfoInterval = 15 * time.Second
)

//################################################# Message ####################

//var MessageNumber int // unique message number (starting from the application launch)

//var AllMessages map[int]MessageData

//type MessageData struct {
//	FieldCount int
//}
