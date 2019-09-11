package pointconfig

type AllCfgData struct {
	Default     CfgFileData
	DefaultJSON CfgFileJSON
	Running     CfgFileData
	RunningJSON CfgFileJSON
}

type CfgFileData map[string]CfgPointData

type CfgFileJSON map[string]JSONPointData

type JSONPointData struct {
	RelIntervalJSON    JSONRelIntervalStruct `json:"RelayOnOffIntervals"`
	TempRelayJSON      JSONTempRelay         `json:"TempRelay"`
	TempRelayArrayJSON []JSONTempRelay       `json:"TempRelayArray"`
}

type JSONData map[string]JSONPointData

type PointCfg struct {
	RelInterv RunRelIntervalStruct
	TempRelay []RunTempRelay
}

type CfgPointData struct {
	List     int      // a field contains bits of available configurations of the point
	Cfg      PointCfg // configuration to use
	CfgSaved PointCfg // saved configuration
}

type AllPointCfgData map[string]CfgPointData
