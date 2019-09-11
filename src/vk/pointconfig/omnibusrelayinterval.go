package pointconfig

import "time"

//=============================================================================
//============== JSON configuration ===========================================
//=============================================================================

type JSONRelInterval struct {
	Gpio     string `json:"Gpio"`
	State    string `json:"State"`
	Interval string `json:"Interval"`
}

type JSONRelIntervalStruct struct {
	Start  JSONRelIntervalArray `json:"Start"`  // array of the point relay default settings (used at the start and exit)
	Base   JSONRelIntervalArray `json:"Base"`   // array of the point relay setting sequences (used between the start and exit)
	Finish JSONRelIntervalArray `json:"Finish"` // array of the point relay setting sequences (used between the start and exit)
}

type JSONRelIntervalArray []JSONRelInterval

/*
type JSONCondition struct {
	MinTemp string `json:"MinTemp"`
	MaxTemp string `json:"MaxTemp"`
	Mask    string `json:"Mask"`
}

type JSONConditions []JSONCondition

type JSONTempRelay struct {
	Conditions JSONConditions `json:"Conditions"`
	Delta      string         `json:"Delta"`
	Fahrenheit string         `json:"Fahrenheit`
	Interval   string         `json:"Interval"`
	Handler    string         `json:"Handler"`
	Gpio       string         `json:"Gpio"`
	State      string         `json:"State"`
	Start      string         `json:"Start"`
	Finish     string         `json:"Finish"`
}
*/

//=============================================================================
//============== Run configuration ===========================================
//=============================================================================

type RunRelInterval struct {
	Gpio    int
	State   int
	Seconds time.Duration
}

type RunRelIntervalArray []RunRelInterval

type RunRelIntervalStruct struct {
	Start  RunRelIntervalArray
	Base   RunRelIntervalArray
	Finish RunRelIntervalArray
}