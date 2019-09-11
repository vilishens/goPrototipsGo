package omnibus

import (
	"fmt"
	"sync"
)

func init() {
	stepList = make(map[string]bool)

	AllMessages = make(map[int]MessageData)

	AllMessages[MsgCdOutputHelloFromStation] = MessageData{FieldCount: MsgPrefixLen + MsgHelloFromStationLen}
	AllMessages[MsgCdInputHelloFromPoint] = MessageData{FieldCount: MsgPrefixLen + MsgHelloFromPointLen}

	// vk-xxx
	AllMessages[MsgCdOutputGpioSet] = MessageData{FieldCount: MsgPrefixLen + MsgSetGpioLen}

	AllMessages[MsgCdInputSuccess] = MessageData{FieldCount: MsgPrefixLen}

	//	LogPointInfo = make(map[int]LogCfg)

	//	LogPointInfo[CfgTypeRelayInterval] = LogCfg{File: "relay-interval", List: []string{LogFileInfo}}
}

func AddStepInList(step string) {
	stepList[step] = true
}

func StepRemoveFromList(step string) {
	delete(stepList, step)
}

func AreStepsInList(steps []string) (err error) {
	for _, v := range steps {
		if _, ok := stepList[v]; !ok {
			err = fmt.Errorf("%q is not in the running step list", v)
			break
		}
	}

	return
}

func StepCount() (count int) {
	return len(stepList)
}

func MessageNumberNext() {

	lock := new(sync.Mutex)
	lock.Lock()
	defer lock.Unlock()

	MessageNumber++
}
