package runrelayinterval

import (
	"net"
	"time"
	vomni "vk/omnibus"
	vcfg "vk/pointconfig"
)

const (
	cmdRestart = 0x0001
)

type AllIndex struct {
	Start  int
	Base   int
	Finish int
}

type RunInterface struct {
	Point string
	State int
	Type  int
	// all point logger  files, key shows bitwise what type of loggers included ("info", "data", ...)
	// The file can have more than one logger (for instance, "info" and "err" info into one file by 2 loggers)
	Logs       []vomni.PointLog
	UDPAddr    net.UDPAddr
	ChErr      chan error
	ChDone     chan int
	ChMsg      chan int
	CfgDefault vcfg.RunRelIntervalStruct
	CfgRun     vcfg.RunRelIntervalStruct
	CfgSaved   vcfg.RunRelIntervalStruct
}

type RunData struct {
	Point string
	State int
	Type  int
	// all point logger  files, key shows bitwise what type of loggers included ("info", "data", ...)
	// The file can have more than one logger (for instance, "info" and "err" info into one file by 2 loggers)
	Logs       []vomni.PointLog
	Index      AllIndex
	UDPAddr    net.UDPAddr
	ChErr      chan error
	ChDone     chan int
	ChMsg      chan string
	ChCmd      chan int
	CfgDefault vcfg.RunRelIntervalStruct
	CfgRun     vcfg.RunRelIntervalStruct
	CfgSaved   vcfg.RunRelIntervalStruct
}

type stage struct {
	once        bool
	runEmptyArr bool
	index       *int
	cfg         vcfg.RunRelIntervalArray
}

type webPoint struct {
	Gpio    string
	State   string
	Seconds string
}

type webPointArr []webPoint

type webPointStruct struct {
	Start  webPointArr
	Base   webPointArr
	Finish webPointArr
}

type RelInterval struct {
	Gpio    int
	State   int
	Seconds time.Duration
}

type RelIntervalArr []RelInterval
