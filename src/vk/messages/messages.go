package messages

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
	vomni "vk/omnibus"
	vparams "vk/params"
	vutils "vk/utils"
)

var MessageList2Send SendMsgArray

func init() {
	MessageList2Send = SendMsgArray{}

	vomni.MessageNumber = 0
}

// GPIO Set
func QeueuGpioSet(point string, addr net.UDPAddr, gpio int, set int) (msg string) {

	data := dataGpioSet(gpio, set)
	nbr := -1
	msg, nbr = dataMsgStr(vomni.MsgCdOutputGpioSet, data)

	msg2SendAdd(point, addr, msg, nbr)

	return
}

func dataGpioSet(gpio int, set int) (data []string) {

	data = append(data, strconv.Itoa(gpio))
	data = append(data, strconv.Itoa(set))

	return data
}

//#######################################################################################

func dataMsgStr(msgCd int, data []string) (msg string, nbr int) {

	vomni.MessageNumberNext()
	nbr = vomni.MessageNumber

	msg = ""
	msg += vparams.Params.StationName + vomni.UDPMessageSeparator
	msg += strconv.Itoa(msgCd) + vomni.UDPMessageSeparator
	msg += strconv.Itoa(nbr)

	for _, v := range data {
		msg += vomni.UDPMessageSeparator + v
	}

	return
}

func msg2SendAdd(point string, addr net.UDPAddr, msg string, nbr int) {

	d := SendMsg{}

	d.PointDst = point
	d.UDPAddr = addr
	d.MessageNbr = nbr
	d.Msg = msg

	MessageList2Send = append(MessageList2Send, d)
}

//#######################################################################################
//#######################################################################################
//#######################################################################################

func Message2SendPlus(addr net.UDPAddr, msgCd int, data []string) {

	msg := message2SendNew(msgCd, data)
	message2SendAdd(addr, msg)
}

func message2SendAdd(addr net.UDPAddr, msg string) {

	d := SendMsg{}

	d.UDPAddr = addr
	d.MessageNbr = vomni.MessageNumber
	d.Msg = msg

	MessageList2Send = append(MessageList2Send, d)
}

func message2SendNew(msgCd int, data []string) (msg string) {

	vomni.MessageNumberNext()

	msg = ""
	msg += vparams.Params.StationName + vomni.UDPMessageSeparator
	msg += strconv.Itoa(msgCd) + vomni.UDPMessageSeparator
	msg += strconv.Itoa(vomni.MessageNumber)

	for _, v := range data {
		msg += vomni.UDPMessageSeparator + v
	}

	return msg
}

func Run(chGoOn chan bool, chDone chan int, chErr chan error) {

	chGoOn <- true
	for {
		time.Sleep(vomni.DelayStepExec)
	}
}

func (d SendMsgArray) MinusIndex(ind int, chDone chan bool) {

	if ind < len(d) {
		lock := new(sync.Mutex)
		lock.Lock()
		defer lock.Unlock()

		MessageList2Send = append(MessageList2Send[:ind], MessageList2Send[ind+1:]...)
	}

	chDone <- true
}

func (d SendMsgArray) MinusNbr(nbr int) {
	ind := -1
	for key, val := range d {
		if val.MessageNbr == nbr {
			ind = key
			break
		}
	}

	if 0 > ind {
		fmt.Printf("Received MSG #%d without record\n", nbr)
		return
	}

	chDone := make(chan bool)
	go d.MinusIndex(ind, chDone)
	<-chDone
}

func TryHello(dst net.UDPAddr, chDone chan bool) {

	msgData := msgStationHello()

	Message2SendPlus(dst, vomni.MsgCdOutputHelloFromStation, msgData)

	fmt.Println("================== try ================================> ", dst, " #", vomni.MessageNumber)

	chDone <- true

}

func msgStationHello() (d []string) {

	_, tzSecs := time.Now().Zone()

	d = make([]string, vomni.MsgHelloFromStationLen)

	d[vomni.MsgIndexHelloFromStationTime] = strconv.Itoa(int(time.Now().Unix()))
	d[vomni.MsgIndexHelloFromStationOffset] = strconv.Itoa(tzSecs)
	d[vomni.MsgIndexHelloFromStationIP] = vparams.Params.IPAddressInternal
	d[vomni.MsgIndexHelloFromStationPort] = strconv.Itoa(vparams.Params.PortUDPInternal)

	return
}

func CheckFieldValue(msg string, index int, value string) (hasValue bool, err error) {

	var msgCd int

	if msgCd, err = msgCode(msg); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	if index >= vomni.AllMessages[msgCd].FieldCount {
		err = vutils.ErrFuncLine(fmt.Errorf("The index %d of the message %#x exceeds the number of fields %d",
			index, msgCd, vomni.AllMessages[msgCd].FieldCount))
		return
	}

	var msgValue string

	if msgValue, err = msgFieldValue(msg, index); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	hasValue = msgValue == value

	return
}

func MessageFields(msg string) (flds []string, err error) {

	if flds, err = msgFields(msg); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	for k, v := range flds {
		newStr := strings.ReplaceAll(v, vomni.UDPMessageSeparator, "")
		if k < (len(flds) - 1) {
			flds[k] = newStr
		}
	}

	return
}

func msgFields(msg string) (flds []string, err error) {

	msgCd, err := msgCode(msg)
	if nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	flds = strings.SplitAfterN(msg, vomni.UDPMessageSeparator, vomni.AllMessages[msgCd].FieldCount)

	return
}

func msgCode(msg string) (cd int, err error) {

	parts := strings.SplitAfterN(msg, vomni.UDPMessageSeparator, vomni.MsgPrefixLen+1)

	cdStr := strings.Replace(parts[vomni.MsgIndexPrefixCd], vomni.UDPMessageSeparator, "", -1)

	if cd, err = strconv.Atoi(cdStr); nil != err {
		return cd, vutils.ErrFuncLine(err)
	}

	if _, has := vomni.AllMessages[cd]; !has {
		return cd, vutils.ErrFuncLine(fmt.Errorf("Not defined Message CD %#x (vk-xxx MSG %q)", cd, msg))
	}

	return
}

func FieldValue(msg string, index int) (value string, err error) {

	return msgFieldValue(msg, index)
}

func msgFieldValue(msg string, index int) (value string, err error) {

	var flds []string

	if flds, err = msgFields(msg); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	value = strings.Replace(flds[index], vomni.UDPMessageSeparator, "", -1)

	return
}

func UpdateField(msg *string, index int, newValue string) (err error) {

	flds, err := msgFields(*msg)
	if nil != err {
		return vutils.ErrFuncLine(err)
	}

	flds[index] = newValue + vomni.UDPMessageSeparator

	*msg = strings.Join(flds, "")

	return
}

func CheckMessageCode(msg string, cd int) (theSame bool, err error) {

	flds, err := msgFields(msg)
	if nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	strCd := strings.ReplaceAll(flds[vomni.MsgIndexPrefixCd], vomni.UDPMessageSeparator, "")

	thisCd, err := strconv.Atoi(strCd)
	if nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	theSame = thisCd == cd

	return
}

func MessageMinusByNbr(nbr int) {
	MessageList2Send.MinusNbr(nbr)
}

func MsgSetGpio(gpio int, set int) {

}
