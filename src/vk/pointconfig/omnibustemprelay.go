package pointconfig

import "time"

//=============================================================================
//============== JSON configuration ===========================================
//=============================================================================

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

//=============================================================================
//============== Run configuration ============================================
//=============================================================================

type RunTempRelay struct {
	Conditions RunConditions
	Delta      float32
	Fahrenheit bool
	Seconds    time.Duration
	Handler    string
	Gpio       int
	State      int
	Start      int
	Finish     int
}

type RunCondition struct {
	MinTemp float32
	MaxTemp float32
	Mask    int
}

type RunConditions []RunCondition
