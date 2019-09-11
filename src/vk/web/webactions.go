package web

import (
	"fmt"
	vscannet "vk/net/netscan"
	vrun "vk/run/pointrun"
)

func rescanWhole() {

	fmt.Println("xxxx\nxxxx\nxxxx\nxxxxx\nxxxx")

	vrun.RescanWhole()
	return
}

func rescanPoint(point string) {

	vrun.RescanPoint(point)
	return

	var lastByte byte //:= vrun.Points[point].Point.UDPAddr.IP[3]

	//	kl := vrun.Points[point].Point.UDPAddr.IP

	baseIP := vrun.Points[point].Point.UDPAddr.IP.To4()

	//_ = err

	fmt.Printf("Defending %v %T\n", baseIP.To4, baseIP)

	if nil != baseIP {
		lastByte = baseIP.To4()[3]
	}

	fmt.Println("###\n###\n###")
	fmt.Println("Kruglaja kapsula", point, "na Marse", vrun.Points[point].Point.UDPAddr.IP, "pod uglom",
		lastByte, "A", baseIP[0], "B", baseIP[1], "C", baseIP[2], "D", baseIP[3])
	fmt.Println("###\n###\n###")

	locDone := make(chan bool)
	locErr := make(chan error)

	go vscannet.IterateIP(locDone, locErr, lastByte, lastByte)

	select {
	case err := <-locErr:

		// chErr <- vutils.ErrFuncLine(err)

		// šeit jāmēģina
		_ = err

	case <-locDone:
		fmt.Printf("Alex Feinsilber --- RESTART")

		// need to wait to be sure all scan messages are done
		// to be able to make all attempts if necessary
		//time.Sleep((vomni.MessageSendRepeatLimit + 1) * vomni.DelaySendMessageRepeat)

		//chGoOn <- true
	}

	fmt.Println("Net problem...")

}
