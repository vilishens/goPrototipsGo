package udp

import (
	"fmt"
	"net"
	"strconv"
	"time"
	vmsg "vk/messages"
	vomni "vk/omnibus"
	vparams "vk/params"
	vpointrun "vk/run/pointrun"
	vutils "vk/utils"
)

func Server(chGoOn chan bool, chDone chan int, chErr chan error) {

	addr := net.UDPAddr{
		Port: vparams.Params.PortUDPInternal,
		IP:   net.ParseIP(vparams.Params.IPAddressInternal),
	}

	conn, err := net.ListenUDP("udp", &addr)

	if err != nil {
		// Something really wrong - let's stop immediately
		addrStr := addr.IP.String() + ":" + strconv.Itoa(addr.Port)
		err = fmt.Errorf("Couldn't get connection of %s --- %v", addrStr, err)
		vutils.LogErr(err)
		chErr <- vutils.ErrFuncLine(err)
		return
	}

	defer conn.Close()

	sendDone := make(chan int)
	sendErr := make(chan error)

	go sendMessages(sendDone, sendErr)

	waitDone := make(chan int)
	waitErr := make(chan error)

	go waitMsg(conn, waitDone, waitErr)

	chGoOn <- true
	select {
	case cd := <-sendDone:
		vutils.LogInfo(fmt.Sprintf("UDP finished with Send RC %d", cd))
	case cd := <-waitDone:
		vutils.LogInfo(fmt.Sprintf("UDP finished with Send RC %d", cd))
	case err := <-sendErr:
		err = fmt.Errorf("Send Error: %s", err)
		vutils.LogErr(err)
		vomni.RootErr <- vutils.ErrFuncLine(err)
	case err := <-waitErr:
		fmt.Errorf("Wait Error: %s", err)
		vutils.LogErr(err)
		vomni.RootErr <- vutils.ErrFuncLine(err)
	}
}

func sendMessages(done chan int, chErr chan error) {

	for {
		time.Sleep(vomni.DelaySendMessage)

		if len(vmsg.MessageList2Send) == 0 {
			// no messages to send
			time.Sleep(vomni.DelaySendMessageListEmpty)
			continue
		}

		for i := 0; i < len(vmsg.MessageList2Send); i++ {
			time.Sleep(vomni.DelaySendMessage)

			if i >= len(vmsg.MessageList2Send) {
				// verify the index isn't out of range
				continue
			}

			// let's make a copy of the message list
			// to avoid the disappered item handling
			copyMsgs := vmsg.MessageList2Send[i]

			chDone := make(chan bool)

			// vk-xxx jāpārraksta tā, lai izmanto kopiju saistībā ar sūtāmo sarakstu
			// (tas vajadzīgs, jo var gadīties, ka apstrādes vidū nodzēš ierakstu,
			// bet turpina meklēt, neesošu ierakstu)

			if "" == copyMsgs.Msg {
				// this is the blank message no need to try to send just remove it
				vutils.LogInfo(fmt.Sprintf("Deleted blank message #%d", copyMsgs.MessageNbr))
				go vmsg.MessageList2Send.MinusIndex(i, chDone)
				<-chDone
				continue
			}

			if !copyMsgs.Last.IsZero() && time.Since(copyMsgs.Last) < vomni.DelaySendMessageRepeat {
				// this is a repeated message but the repeat interval isn't passed yet
				continue
			}

			// vk-xxx jāpieliek Msg Repeat caur objektu
			vmsg.MessageList2Send[i].Repeat++

			// Need to update the station time if it is the repeated message and 'Hello From Station'
			if err := updateMessageStationTime(i); nil != err {

				fmt.Println("vk-xxx ===================> Veniamin Reshetnikoff")

				chErr <- err
				return
			}

			if copyMsgs.Repeat > vomni.MessageSendRepeatLimit {

				vutils.LogInfo(fmt.Sprintf("Deleted message #%d due to the exceeded send repeat limit", copyMsgs.MessageNbr))

				// set the point (if it has signed in) as disconnected
				vpointrun.SetDisconnectedPoint(vmsg.MessageList2Send[i].UDPAddr)

				go vmsg.MessageList2Send.MinusIndex(i, chDone)

				<-chDone
				continue
			}

			vmsg.MessageList2Send[i].Last = time.Now()

			if err := SendToAddress(vmsg.MessageList2Send[i].UDPAddr, vmsg.MessageList2Send[i].Msg); nil != err {
				// write the error in log
				vutils.LogErr(err)
			}
		}
	}
}

// Need to update the station time if it is the repeated message and 'Hello From Station'
func updateMessageStationTime(i int) (err error) {

	if 1 < vmsg.MessageList2Send[i].Repeat {
		// it is the repeated message

		var isStationHello bool
		isStationHello, err = vmsg.CheckMessageCode(vmsg.MessageList2Send[i].Msg, vomni.MsgCdOutputHelloFromStation)

		if nil != err {
			return
		}

		if isStationHello {
			// this message is 'Hello From Station'
			strTime := strconv.Itoa(int(time.Now().Unix()))

			vmsg.UpdateField(&vmsg.MessageList2Send[i].Msg, vomni.MsgPrefixLen+vomni.MsgIndexHelloFromStationTime, strTime)
		}
	}

	return
}

func SendToAddress(addr net.UDPAddr, msg string) (err error) {

	addrStr := addr.IP.String() + ":" + strconv.Itoa(addr.Port)

	chErr := make(chan error)
	go sendToAddress(addrStr, msg, chErr)

	select {
	case err = <-chErr:
		return
	}

	return <-chErr
}

func sendToAddress(addr string, msg string, chErr chan error) (err error) {

	conn, err := net.Dial("udp", addr)
	if err != nil {
		err = vutils.ErrFuncLine(fmt.Errorf("Connection ERROR: %v", err))
		chErr <- err
		return
	}

	if nil == err {
		defer conn.Close()

		if _, err = conn.Write([]byte(msg)); nil != err {
			err = vutils.ErrFuncLine(fmt.Errorf("SentToAddress ERROR: %v", err))
		}
	}

	chErr <- err
	return
}

func waitMsg(conn *net.UDPConn, done chan int, chErr chan error) {

	for {
		// waiting, waiting, ... UDP
		time.Sleep(vomni.DelayWaitMessage)

		buff := make([]byte, 4096)
		nn, msgAddr, err := conn.ReadFromUDP(buff)
		if err != nil {
			continue
		}

		if len(buff) == 0 {
			continue
		}

		msg := string(buff[:nn])

		locErr := make(chan error)
		go vpointrun.MessageReceived(msg, locErr)

		if err = <-locErr; nil != err {
			chErr <- vutils.ErrFuncLine(err)
			return
		}

		if err != nil {
			vutils.LogErr(fmt.Errorf("The received message %q (address %s:%d) %error %q", msg,
				msgAddr.IP.String(), msgAddr.Port, err.Error()))
		}

		buff = []byte{}
	}
}
