package netscan

import (
	"fmt"
	"net"
	"time"

	vmsg "vk/messages"
	vomni "vk/omnibus"
	vparams "vk/params"
	vutils "vk/utils"
)

// ScanOctet - goes through x.x.x.Start till x.x.x.End IP addresses and sends Hi message to points at these addresses
func ScanOctet(chGoOn chan bool, chDone chan int, chErr chan error) {

	locDone := make(chan bool)
	locErr := make(chan error)

	go IterateIP(locDone, locErr, IPStart, IPEnd)

	select {
	case err := <-locErr:
		chErr <- vutils.ErrFuncLine(err)
	case <-locDone:
		fmt.Println("Alex Feinsilber")

		// need to wait to be sure all scan messages are done
		// to be able to make all attempts if necessary
		//time.Sleep((vomni.MessageSendRepeatLimit + 1) * vomni.DelaySendMessageRepeat)

		chGoOn <- true
	}
}

func IterateIP(chDone chan bool, chErr chan error, start byte, end byte) {

	baseIP := net.ParseIP(vparams.Params.IPAddressInternal).To4()
	if nil == baseIP {
		chErr <- vutils.ErrFuncLine(fmt.Errorf("The internal IP address is not defined yet"))
		return
	}

	for i := start; i <= end; i++ {

		ip := net.IP{baseIP[0], baseIP[1], baseIP[2], byte(i)}

		if ip.String() == vparams.Params.IPAddressInternal {
			// don't send to the station itself (the same IP address)
			continue
		}

		locErr := make(chan error)
		locDone := make(chan bool)

		if i == IPEnd {
			fmt.Printf("\n\t\t\t\t******** END IP %q *******\n", ip.String())
		}

		go tryPointConn(ip, locDone, locErr)

		select {
		case <-locDone:
			//
		case err := <-locErr:
			vutils.LogErr(fmt.Errorf("Error %q received from IP %s:%d Hello", err.Error(),
				ip.String(), vparams.Params.PortUDPInternal))
		}

		time.Sleep(vomni.DelayBetweenIPHello)
	}

	fmt.Println("#################################################### Scan NET done...")
	for k, v := range vmsg.MessageList2Send {
		fmt.Printf("PEC VISIEM SCANIEM %2d --- IP %q NBR %d\n", k, v.UDPAddr.IP.String(), v.MessageNbr)
	}

	// need to wait to be sure all scan messages are done
	// to be able to make all attempts if necessary
	time.Sleep((vomni.MessageSendRepeatLimit + 1) * vomni.DelaySendMessageRepeat)

	chDone <- true
}

func tryPointConn(ip net.IP, chDone chan bool, chErr chan error) {

	dstAddr := net.UDPAddr{IP: ip, Port: vparams.Params.PortUDPInternal}

	locDone := make(chan bool)
	locErr := make(chan error)

	go vmsg.TryHello(dstAddr, locDone)

	select {
	case err := <-locErr:
		vutils.LogErr(fmt.Errorf("Hello try failed: IP %s:%d", dstAddr.IP.String(), dstAddr.Port))
		chErr <- err
		break
	case <-locDone:
		chDone <- true
		break
	}
}
