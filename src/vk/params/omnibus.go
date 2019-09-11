package params

type ParamData struct {
	StationName string

	LogMainPath  string
	LogPointPath string

	PortUDPInternal int
	PortSSHInternal int
	PortWEBInternal int
	PortSSHExternal int
	PortWEBExternal int

	RotateMainTmpl       string
	RotatePointDataTmpl  string
	RotatePointInfoTmpl  string
	RotateRunCfg         string
	RotateRunSecs        int
	RotateStatusFileName string

	WebStaticPrefix string
	WebStaticDir    string
	WebTemplateDir  string

	IPAddressInternal string
	IPAddressExternal string

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
