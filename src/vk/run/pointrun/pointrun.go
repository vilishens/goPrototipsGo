package pointrun

import (
	"fmt"
	"net"
	"strconv"
	"time"
	vmsg "vk/messages"
	vnetscan "vk/net/netscan"
	vomni "vk/omnibus"
	vrunrelayinterval "vk/run/relayinterval"
	vutils "vk/utils"
)

var Points map[string]*PointRun
var listSigned map[string]net.UDPAddr

func init() {
	Points = make(map[string]*PointRun)
	listSigned = make(map[string]net.UDPAddr)
}

func Runners() {
	relayIntervalRunners()
}

func relayIntervalRunners() {
	for k, v := range vrunrelayinterval.RunningPoints {

		if _, has := Points[k]; !has {
			// it is required to create a new point running object from the template
			addNewPointRun(k)
		}

		// set the type of the Point
		tPoint := Points[k].Point
		tPoint.Type |= v.Type

		// save the the configuration data
		tRun := Points[k].Run
		tRun[v.Type] = v

		// put all data into the point running object
		Points[k].Point = tPoint
		Points[k].Run = tRun
	}
}

func StopAll(chDone chan bool) {

	fmt.Println("????????????????????????????????????????????????????????")
	fmt.Println("????????????????????????????????????????????????????????")
	fmt.Println("????????????????????????????????????????????????????????")
	fmt.Println("????????????????????????????????????????????????????????")
	fmt.Println("????????????????????????????????????????????????????????")
	fmt.Println("????????????????????????????????????????????????????????")
	fmt.Println("????????????????????????????????????????????????????????")
	fmt.Println("????????????????????????????????????????????????????????")
	fmt.Println("????????????????????????????????????????????????????????")
	fmt.Println("????????????????????????????????????????????????????????")
	fmt.Println("????????????????????????????????????????????????????????")

	ind := len(vomni.CfgListSequence) - 1

	for ind = len(vomni.CfgListSequence) - 1; ind >= 0; ind-- {
		// start with the last item in the point configuration sequence
		// (opposite direction to the start of configuration)

		cfgCd := vomni.CfgListSequence[ind]

		for k, v := range Points {
			if 0 < v.Point.Type&cfgCd {
				// the point has this configuration

				if (0 == (v.Point.State & vomni.PointStateDisconnected)) &&
					(0 < (v.Point.State & vomni.PointStateSigned)) {
					// this point is connected and the connection ir active
					Points[k].setState(vomni.PointStateStoppingNow, true)

					fmt.Println("########################################################")
					fmt.Println("!!!!!\n!!!!!\n", v.Point.Point, "\n!!!!!\n!!!!!", k)
					fmt.Println("########################################################")

					v.Run[cfgCd].Cmd(vomni.PointCmdStopCfg)

				}
			}
		}

	}

	pCount := runningTotal()
	for pCount > 0 {
		time.Sleep(500 * time.Millisecond)
		pCount = runningTotal()
	}

	chDone <- true
}

func runningTotal() (count int) {

	for ind := len(vomni.CfgListSequence) - 1; ind < len(vomni.CfgListSequence); ind++ {

		cfgCd := vomni.CfgListSequence[ind]
		for _, v := range Points {
			if 0 < v.Point.Type&cfgCd {
				count += v.Run[cfgCd].GetRunTotal()
				break
			}
		}
	}

	return count
}

func RunStart(chGoOn chan bool, chDone chan int, chErr chan error) {

	locGoOn := make(chan bool)
	locDone := make(chan int)
	locErr := make(chan error)

	go scanOctetAddrs(vnetscan.IPStart, vnetscan.IPEnd, locGoOn, locDone, locErr)

	stop := false
	for {
		select {
		case err := <-locErr:
			chErr <- err
			return
		case done := <-chDone:
			chDone <- done
			return
		case <-locGoOn:
			fmt.Println("### Kurtenkov ###")
			stop = true
		}

		if stop {
			break
		}
	}

	fmt.Println("Alex Sitkovetsky ")
	chGoOn <- true

	fmt.Println("TAGAD starts", len(listSigned))
	fmt.Printf("TAGAD oooooo %+v\n", listSigned)

	for {
		time.Sleep(vomni.DelayStepExec)
	}
}

func RescanPoint(point string) {

	baseIP := Points[point].Point.UDPAddr.IP.To4()

	if nil == baseIP {
		err := fmt.Errorf("Can't rescan the point %q due to abscense or invalid of IP address", point)
		vutils.LogErr(err)
	}

	lastByte := baseIP.To4()[3]

	locGoOn := make(chan bool)
	locErr := make(chan error)
	locDone := make(chan int)

	go scanOctetAddrs(lastByte, lastByte, locGoOn, locDone, locErr)
	select {
	case err := <-locErr:
		vutils.LogErr(err)
		return
	case <-locDone:
		fmt.Printf("Alex Feinsilber --- RESTART SIGN--> šeit")
		return
	case <-locGoOn:
		fmt.Printf("Alex Feinsilber --- RESTART GO ON--> šeit")
		return
	}
}

func RescanWhole() {

	fmt.Println("@@@@\n@@@@\n@@@@@\n@@@@@\n@@@@@")

	locGoOn := make(chan bool)
	locErr := make(chan error)
	locDone := make(chan int)

	go scanOctetAddrs(vnetscan.IPStart, vnetscan.IPEnd, locGoOn, locDone, locErr)
	select {
	case err := <-locErr:
		vutils.LogErr(err)
		return
	case <-locDone:
		fmt.Printf("Alex Feinsilber --- RESTART SIGN--> šeit")
		return
	case <-locGoOn:
		fmt.Printf("Alex Feinsilber --- RESTART GO ON--> šeit")
		return
	}
}

func scanOctetAddrs(first byte, last byte, chGoOn chan bool, chDone chan int, chErr chan error) {

	// prepare storage for signed in points
	listSigned = map[string]net.UDPAddr{}

	locGoOn := make(chan bool)
	locErr := make(chan error)
	locDone := make(chan int)

	fmt.Println("RESCAN POINT >>>>>")
	fmt.Println("RESCAN POINT >>>>>")
	fmt.Printf("RESCAN POINT >>>>> F %d L %d\n", first, last)
	fmt.Println("RESCAN POINT >>>>>")
	fmt.Println("RESCAN POINT >>>>>")

	go vnetscan.IterateIP(locGoOn, locErr, first, last)

	select {
	case err := <-locErr:
		chErr <- vutils.ErrFuncLine(err)

		//vutils.LogErr(err)
		//Points[point].Run.LogStr(vomni.LogFileCdErr, err.String())
		return
	case <-locGoOn:
		chGoOn <- true
		fmt.Printf("Alex Feinsilber --- RESCAN --> DONE\n")
	}

	go startSigned(locGoOn, locDone, locErr)

	select {
	case err := <-locErr:
		vutils.LogErr(err)
		return
	case <-locDone:
		fmt.Printf("Alex Feinsilber --- RESTART SIGN--> šeit")
	case <-locGoOn:
		fmt.Printf("Alex Feinsilber --- RESTART GO ON--> šeit")
	}

}

//#######################################################################

func startSigned(chGoOn chan bool, chDone chan int, chErr chan error) {

	listHandled := make(map[string]bool) // list of signed already and handled points
	//	listStart := make(map[string]bool)   // list of points in start state

	for k, v := range listSigned {
		fmt.Printf("POINT %q ADDR %+v LEN %d\n", k, v, len(listSigned))
	}

	// Vajad izmainīt
	// vispirms jāapstrādā pats punkts, pec tam jaapstrādā tā konfigurācijas

	for _, cfgType := range vomni.CfgListSequence {
		// start all point configuration, sequence set in startSequence
		// Sequence can be important some times (for instance, to check the point ready state)
		for point, addr := range listSigned {
			logStr := ""
			pData := new(PointRun)
			hasCfg := false

			err := error(nil)

			//delete(listSigned, point)

			fmt.Println("Did I missed???")
			fmt.Println("Did I missed???")
			fmt.Printf("Time 2 sign %q\n", point)
			fmt.Println("Did I missed???")
			fmt.Println("Did I missed???")

			// check if the point has its object
			if pData, err = checkStartPointExistence(point, addr); nil == err {
				// save the point address
				pData.setUDPAddr(addr)
				// check the point and configuration existence
				hasCfg, err = pData.checkStartCfgExistence(cfgType)
			}

			if nil != err {
				// the error was found
				vutils.LogErr(err)
				chErr <- vutils.ErrFuncLine(err)
				return
			}

			if !hasCfg {
				// there is no configuration of this type
				continue
			}

			pointStart := false
			cfgStart := false

			fmt.Println("Point >>>", point, "<<< has type", cfgType)

			// the point has configuration of this type
			if _, has := listHandled[point]; !has {
				// the point isn't handled yet
				fmt.Println("Point >>>", point, "<<< hasn't handled yet")

				pointStart, logStr = pData.handlePointStart()

				// put messages about signed in into log
				vutils.LogInfo(logStr)

				pData.setState(vomni.PointStateSigned, true)
				if pointStart {
					pData.setState(vomni.PointStateDisconnected, false)
					pData.setState(vomni.PointStateFrozen, false)
				}

				listHandled[point] = true
				fmt.Println("Point >>>", point, "<<< has handled now")
			}

			pCfg := pData.Run[cfgType]

			if !pCfg.Ready() {
				// handle non Ready point configuration
				errStr := pData.handleNonReadyConfiguration(cfgType)

				if errStr != "" {
					vutils.LogErr(fmt.Errorf("%s", errStr))

					// send log to the point configuration
					// (it succeeds only if the point configuration was ready (rotate files have been started))
					pCfg.LogStr(vomni.LogFileCdErr, errStr)
				}

				pCfg.SetState(vomni.PointCfgStateReady, false)
				pCfg.SetState(vomni.PointCfgStateUnavailable, false)

				continue
			} else {
				if cfgStart, logStr, err = pData.handleReadyConfiguration(cfgType); nil != err {
					chErr <- err
					return
				}

				if logStr != "" {
					vutils.LogInfo(logStr)

					// send log to the point configuration
					// (it succeeds only if the point configuration was ready (rotate files were started))
					pCfg.LogStr(vomni.LogFileCdInfo, logStr)
				}

				// remember this configuration state
				pCfg.SetState(vomni.PointCfgStateUnavailable, false)
				pCfg.SetState(vomni.PointCfgStateReady, true)
			}

			if pointStart || cfgStart {
				locGoOn := make(chan bool)
				locDone := make(chan int)
				locErr := make(chan error)

				pCfg.SetUDPAddr(addr)

				go pCfg.LetsGo(locGoOn, locDone, locErr)

				select {
				case <-locGoOn:
				case cd := <-locDone:

					if (cd & vomni.PointCmdExitCfg) > 0 {
						cfgType = cd & vomni.PointCmdOptionBits

						delete(Points[point].Run, cfgType)
						if 0 == len(Points[point].Run) {
							delete(Points, point)
						}
					}
				case err := <-locErr:
					chErr <- err
				}
			}
		}
	}
}

func scanNet(chGoOn chan bool, chDone chan int, chErr chan error) {

	locGoOn := make(chan bool)
	locDone := make(chan int)
	locErr := make(chan error)

	go vnetscan.ScanOctet(locGoOn, locDone, locErr)

	for {
		select {
		case err := <-locErr:
			chErr <- err
			return
		case done := <-chDone:
			chDone <- done
			return
		case <-locGoOn:
			chGoOn <- true
			return
		}
	}
}

func MessageReceived(msg string, chErr chan error) {

	fmt.Println("vk-xxx @@@@@@ SITKOVETSKY @@@@@ MSG", msg)

	var err error
	var flds []string
	if flds, err = vmsg.MessageFields(msg); nil != err {
		chErr <- vutils.ErrFuncLine(err)
		return
	}

	fmt.Printf("SITKOVETSKY %+q\n", flds)

	msgNbr, err := strconv.Atoi(flds[vomni.MsgIndexPrefixNbr])
	if nil != err {
		vutils.LogErr(fmt.Errorf("The Msg Number error of Msg %q", msg))
		chErr <- vutils.ErrFuncLine(err)
		return
	}

	locErr := make(chan error)
	locDelete := make(chan bool)

	go messageReceived(flds, locDelete, locErr)
	select {
	case <-locDelete:
		vmsg.MessageList2Send.MinusNbr(msgNbr)
	case err = <-locErr:
		break
	}

	chErr <- err
}

func messageReceived(flds []string, chDelete chan bool, chErr chan error) {

	var err error
	msgCd := -1

	if msgCd, err = strconv.Atoi(flds[vomni.MsgIndexPrefixCd]); nil != err {
		err = fmt.Errorf("The Msg Code error of Msg %v", flds)
		vutils.LogErr(err)
		chErr <- vutils.ErrFuncLine(err)
	}

	locErr := make(chan error)
	locDone := make(chan bool)
	locDelete := make(chan bool)

	switch msgCd {
	case vomni.MsgCdInputHelloFromPoint:
		fmt.Println("..............................................................")
		fmt.Printf("........................ RUNNING HELLO! %s\n", flds[vomni.MsgIndexPrefixSender])
		fmt.Println("..............................................................")

		//go handleHelloFromPoint(flds, locDone, locErr)

		go addSignIn(flds, locDelete, locErr)

	case vomni.MsgCdOutputHelloFromStation:
		// this is the hello message from another station
		// just ignore it and send delete it signal
		chDelete <- true
		return
	case vomni.MsgCdInputSuccess:
		// don't do anything - just send delete it signal

		fmt.Println(flds[vomni.MsgIndexPrefixSender], "@@@@@@@@@@@@@ vk-xxx -------> SUCCESS received")

		chDelete <- true
		return
	default:
		chErr <- vutils.ErrFuncLine(fmt.Errorf("RECEIVED->RECEIVED->RECEIVED unknowm CMD %d", msgCd))
		fmt.Println("Eduards")
		return
	}

	select {
	case <-locDone:
		// the done code received
	case <-locDelete:

		fmt.Println("vk-xxx Colombus")

		chDelete <- true
	case err = <-locErr:
		// the error received
		vomni.RootErr <- err
		return
	}

	chErr <- err
}

func addSignIn(flds []string, chDelete chan bool, chErr chan error) {

	fmt.Println(".................................................>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	point := flds[vomni.MsgIndexPrefixSender]

	fmt.Println("LIHTENBERG-lihtengerg-LIHTENBERG--> jāieliek normālā konfigurācijas pārbaude")

	if _, has := Points[point]; !has {

		//		fmt.Println("LIHTENBERG-lihtengerg-LIHTENBERG--> jāieliek normālā konfigurācijas pārbaude")

		return
	}

	addr, ok := getUDPAddr(flds, vomni.MsgPrefixLen+vomni.MsgIndexHelloFromPointIP, vomni.MsgPrefixLen+vomni.MsgIndexHelloFromPointPort)

	if ok {
		listSigned[point] = addr
	}

	// send back the flag to delete this message
	chDelete <- true

	//	fmt.Printf("vk-xxx PEVICHKA! %+v\nPoint %q\n", flds, point)

	fmt.Printf("vk-xxx PEVICHKA! %+v\nPoint %q UDP %+v\n", flds, Points[point].Point.Point, Points[point].Point.UDPAddr)
}

func SetDisconnectedPoint(addr net.UDPAddr) (point string) {
	for k, v := range Points {
		if vutils.Equal(addr, v.Point.UDPAddr) &&
			(0 != v.Point.State&vomni.PointStateSigned) &&
			(0 == v.Point.State&vomni.PointStateDisconnected) {

			Points[k].setState(vomni.PointStateDisconnected, true)
			str := fmt.Sprintf("Point %q lost connection", k)

			vutils.LogInfo(str)

			doneCd := vomni.DoneDisconnected // disconnect regular
			if 0 < (Points[k].Point.Type & vomni.PointStateStoppingNow) {
				doneCd = vomni.DoneExit // disconnect while stopping
			}

			// send disconnection code to all configurations of the point
			for _, v := range Points[k].Run {
				v.GetDone(doneCd)
			}
		}
	}

	return
}

func getUDPAddr(flds []string, ipInd int, portInd int) (addr net.UDPAddr, ok bool) {

	intPort, err := strconv.Atoi(flds[portInd])
	if nil != err {
		err = vutils.ErrFuncLine(fmt.Errorf("A message received (%v) with the wrong Port format %q - %s",
			flds,
			flds[portInd],
			err.Error()))
		vutils.LogErr(err)
	}

	netIP := net.ParseIP(flds[ipInd])
	if nil == netIP {
		err = vutils.ErrFuncLine(fmt.Errorf("A message received (%v) with the invalid IP %q",
			flds,
			flds[ipInd]))
		vutils.LogErr(err)
	}

	if nil != err {
		return
	}

	addr = net.UDPAddr{IP: netIP, Port: intPort}

	return addr, true
}

func (d *PointRun) setState(state int, on bool) {
	if on {
		d.Point.State |= state
	} else {
		d.Point.State &^= state
	}
}

func (d *PointRun) setUDPAddr(addr net.UDPAddr) {
	d.Point.UDPAddr = addr
}

func addNewPointRun(point string) {
	newP := new(PointRun)

	newP.Point.Point = point
	newP.Point.State = vomni.PointStateUnknown
	newP.Point.Type = vomni.CfgTypeUnknown
	newP.Point.UDPAddr = net.UDPAddr{}

	newP.Run = make(map[int]Runner)

	Points[point] = newP
}

func (d *PointRun) handlePointStart() (start bool, str string) {

	if 0 != (d.Point.State & vomni.PointStateDisconnected) {
		// this point was signed in, but later disconnected
		// need to restart again
		str = fmt.Sprintf("START SIGNED *** Point %q signed in AGAIN", d.Point.Point)
		start = true // need to restart
	} else if 0 == (d.Point.State & vomni.PointStateSigned) {
		// the point wasn't signed in, need to start from scratch
		str = fmt.Sprintf("START SIGNED *** Point %q signed in", d.Point.Point)
		start = true
	} else {
		// the point was signed and not disconnected, to update the address is enough
		str = fmt.Sprintf("START SIGNED *** Point %q (signed in already) saves the new UDP address %s:%d",
			d.Point.Point, d.Point.UDPAddr.IP.String(), d.Point.UDPAddr.Port)
	}

	// put messages about signed in into log
	//		vutils.LogInfo(logStr)

	// set the clean signed state
	//		pData.setState(vomni.PointStateDisconnected, false)
	//		pData.setState(vomni.PointStateSigned, true)

	//		listHandled[point] = true

	return
}

func (d *PointRun) handleCfgStart(cfg int) (start bool, str string) {
	return
}

func checkStartPointExistence(point string, addr net.UDPAddr) (pt *PointRun, err error) {

	has := false

	if pt, has = Points[point]; !has {
		// The point with no configuration signed
		// It is wrong so shouldn't be
		err = fmt.Errorf("The point %q (%v) sent SignIn message, but there is no configuration of this point",
			point, addr)
	}

	return
}

func (d *PointRun) checkStartCfgExistence(cfg int) (has bool, err error) {

	if 0 != (cfg & d.Point.Type) {
		if _, has = d.Run[cfg]; !has {
			// The point should have this configuration but doesn't
			// It is wrong so shouldn't be
			err = fmt.Errorf("The point %q (%v) should have %q configuration but DOESN'T",
				d.Point.Point, d.Point.UDPAddr, vomni.PointCfgData[cfg].CfgStr)
		}
	}

	return
}

func (d *PointRun) handleNonReadyConfiguration(cfg int) (errStr string) {

	strState := ""
	pCfg := d.Run[cfg]
	state := pCfg.GetState()

	if 0 != (vomni.PointCfgStateUnavailable & state) {
		// this configuration has been unavailable already
		// no log message required
	} else if vomni.PointCfgStateUnknown == state {
		strState = "not started yet"
	} else if 0 != (vomni.PointCfgStateReady & state) {
		strState = "was ready"
	}

	if strState != "" {
		errStr = fmt.Sprintf("The point %q (%s) configuration %q is not ready",
			d.Point.Point,
			strState,
			vomni.PointCfgData[cfg].CfgStr)
	}

	return
}

func (d *PointRun) handleReadyConfiguration(cfg int) (start bool, logStr string, err error) {

	strState := ""
	pCfg := d.Run[cfg]
	state := pCfg.GetState()

	if vomni.PointCfgStateUnknown == state {
		// this the very 1st start of the configuration

		// start rotation of the log files
		if err = pCfg.StartRotate(); nil != err {
			return
		}

		// this the very 1st start of the configuration
		start = true
		strState = "wasn't started yet"
	} else if 0 != (vomni.PointCfgStateUnavailable & state) {
		start = true
		strState = "was unavailable"
	}

	if strState != "" {
		logStr = fmt.Sprintf("The point %q (%s) configuration %q is ready",
			d.Point.Point,
			strState,
			vomni.PointCfgData[cfg].CfgStr)
	}

	return
}
