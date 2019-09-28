package omnibus

import (
	"log"
	"os"
)

// the point log data file configuration
type PointLog struct {
	LogFile    string              // the full path of the data file
	LogTmpl    string              // the full path of the rotate configuration template file
	LogFilePtr *os.File            // the pointer to the opened data file
	Loggers    map[int]PointLogger // all loggers linked to the data file with the key of the logger bitwise code
}

// the logger configuration
type PointLogger struct {
	LogPrefix string      // the prefix
	Logger    *log.Logger // the logger
}

type LogPointData struct {
	LogCd     int
	FileEnd   string
	LogPrefix string
}

type CfgPointData struct {
	CfgCd  int
	CfgStr string
}

type MessageData struct {
	FieldCount int
}

type WebAllPointData struct {
	List []string
	Data map[string]WebPointData
}

type CfgPlusData struct {
	Name string
}

type WebPointData struct {
	Point        string
	State        int
	Type         int
	Signed       bool
	Disconnected bool
	Frozen       bool
	CfgList      []int
	CfgInfo      map[int]CfgPlusData
	CfgRun       map[int]interface{}
	CfgSaved     map[int]interface{}
	CfgDefault   map[int]interface{}
	CfgIndex     map[int]interface{}
	CfgState     map[int]interface{}
}

//#################################################################
//#################################################################
//#################################################################

// vai šitas vajadzīgs????
type LogCfg struct {
	File string
	List []string
}

//??????????????????????
var LogPointInfo map[int]LogCfg
