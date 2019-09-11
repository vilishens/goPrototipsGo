package omnibus

const (
	MsgCdOutputHelloFromStation = 0x00000001 // Output <station name><MsgCdOutputHelloFromStation><msgNbr><station UTC seconds><station time offset><stationIP><stationPort>
	MsgCdInputHelloFromPoint    = 0x00000002 // Input  <point name><MsgCdInputHelloFromPoint><msgNbr><pointIP><pointPort>
	MsgCdOutputGpioSet          = 0x00000004 // Output <station name><MsgCdOutputGpioSet><msgNbr><Gpio><set value>
	MsgCdInputSuccess           = 0x00000008 // Input  <point name><MsgCdInputSuccess><msgNbr>
)

const (
	UDPMessageSeparator = ":::"

	MsgIndexPrefixSender = 0
	MsgIndexPrefixCd     = 1
	MsgIndexPrefixNbr    = 2

	MsgPrefixLen = 3
)

// Hello From Station
const (
	MsgIndexHelloFromStationTime   = 0
	MsgIndexHelloFromStationOffset = 1
	MsgIndexHelloFromStationIP     = 2
	MsgIndexHelloFromStationPort   = 3

	MsgHelloFromStationLen = 4
)

// Hello From Point
const (
	MsgIndexHelloFromPointIP   = 0
	MsgIndexHelloFromPointPort = 1

	MsgHelloFromPointLen = 2
)

// Set Gpio
const (
	MsgIndexSetGpioGpio = 0
	MsgIndexSetGpioSet  = 1

	MsgSetGpioLen = 2
)
