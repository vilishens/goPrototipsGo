package cfg

type CfgData struct {
	StationName string `json:"StationName"`

	RotateMainTmpl       string `json:"RotateMainTmpl"`
	RotatePointDataTmpl  string `json:"RotatePointDataTmpl"`
	RotatePointInfoTmpl  string `json:"RotatePointInfoTmpl"`
	RotateRunCfg         string `json:"RotateRunCfg"`
	RotateRunSecs        string `json:"RotateRunIntervalSecs"`
	RotateStatusFileName string `json:"RotateStatusFileName"`

	LogPointPath string `json:"LogPointPath"`

	PortUDPInternal string `json:"PortUDPInternal"`
	PortSSHInternal string `json:"PortSSHInternal"`
	PortWEBInternal string `json:"PortWEBInternal"`
	PortSSHExternal string `json:"PortSSHExternal"`
	PortWEBExternal string `json:"PortWEBExternal"`

	WebStaticPrefix string `json:"WEBStaticPrefix"`
	WebStaticDir    string `json:"WEBStaticDir"`
	WebTemplateDir  string `json:"WEBTemplateDir"`

	IPExternalAddressCmds  []string `json:"IPExternalAddressCmds"`
	NetExternalRequirement string   `json:"NetExternalRequired"`

	PointConfigDefaultFile string `json:"PointConfigDefaultFile"`
	PointConfigFile        string `json:"PointConfigFile"`

	MessageEmailAddress string `json:"MessageEmailAddress"`
	MessageSMTPHost     string `json:"MessageSMTPHost"`
	MessageSMTPUser     string `json:"MessageSMTPUser"`
	MessageSMTPPass     string `json:"MessageSMTPPass"`
	MessageSMTPPort     string `json:"MessageSMTPPort"`
}

type CfgFinalData struct {
	StationName string
	LogMainPath string

	RotateMainTmpl       string
	RotatePointDataTmpl  string
	RotatePointInfoTmpl  string
	RotateRunCfg         string
	RotateRunSecs        int
	RotateStatusFileName string

	LogPointPath string

	PortUDPInternal int
	PortSSHInternal int
	PortWEBInternal int
	PortSSHExternal int
	PortWEBExternal int

	WebStaticPrefix string
	WebStaticDir    string
	WebTemplateDir  string

	IPExternalAddressCmds  []string
	NetExternalRequirement int

	PointConfigDefaultFile string
	PointConfigFile        string

	MessageEmailAddress string
	MessageSMTPHost     string
	MessageSMTPUser     string
	MessageSMTPPass     string
	MessageSMTPPort     int
}
