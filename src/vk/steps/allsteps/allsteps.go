package allsteps

import (
	"fmt"

	//	vcfg "vk/cfg"
	vomni "vk/omnibus"
	vutils "vk/utils"

	vstep "vk/steps/step"

	//	schecknet "vk/steps/stepchecknet"
	scfg "vk/steps/stepconfig"
	sparams "vk/steps/stepparams"

	//	spointcfg "vk/steps/steppointconfig"
	//	spointscan "vk/steps/steppointscan"
	//	srunpoints "vk/steps/steprunpoints"
	smsg "vk/steps/stepmessages"
	snetinfo "vk/steps/stepnet/stepnetinfo"
	spointcfg "vk/steps/steppointconfig"
	spointready "vk/steps/steppointready"
	spointrun "vk/steps/steppointrun"
	srotatemain "vk/steps/steprotatemain"
	sstart "vk/steps/stepstart"
	sudp "vk/steps/stepudp"
	sweb "vk/steps/stepweb"
)

var steps = make(map[string]vstep.Step)
var stepSequence []string

func init() {
	initSteps()
}

func initSteps() {
	addStep(&(sstart.ThisStep))      // the very first routines: CLI flags, ...
	addStep(&(scfg.ThisStep))        // application configuration
	addStep(&(sparams.ThisStep))     // prepare application configuration as parameters
	addStep(&(snetinfo.ThisStep))    // get and check frequently net info, send email about it state if necessary
	addStep(&(srotatemain.ThisStep)) // set rotation of the main (application) log file
	addStep(&(sweb.ThisStep))        // start WEB server
	addStep(&(smsg.ThisStep))        // messages
	addStep(&(spointcfg.ThisStep))   // point configuration
	addStep(&(spointready.ThisStep)) // point configuration preparing to run
	addStep(&(sudp.ThisStep))        // starts UDP
	addStep(&(spointrun.ThisStep))   // runpoints
}

func addStep(s vstep.Step) {
	sName := s.StepName()

	if _, exists := steps[sName]; exists {
		panic(fmt.Sprintf("ALERT! Step '%s' used more than once (ind %d)", sName, len(stepSequence)))
	}

	stepSequence = append(stepSequence, sName)
	steps[sName] = s
}

func DoSteps(chDone chan int) {

	locDone := make(chan int)

	go doAllSteps(locDone)

	done := <-locDone

	chDone <- done
}

func doAllSteps(chanDone chan int) {

	chErr := make(chan error)
	chDone := make(chan int)
	chGoOn := make(chan bool)

	stop := false
	err := error(nil)
	done := 0

	for _, s := range stepSequence {
		thisS := steps[s]

		str := fmt.Sprintf("===== Step %q -> started", thisS.StepName())
		fmt.Println(str)
		vutils.LogStr(vomni.LogInfo, fmt.Sprintf(str))
		go vstep.Execute(thisS, chDone, chGoOn, chErr)

		select {
		case <-chGoOn:
			vomni.AddStepInList(thisS.StepName())

		case err = <-chErr:
			stop = true

		case err = <-vomni.StepErr:
			stop = true

		case done = <-chDone:
			stop = true

		case done = <-vomni.RootDone:
			stop = true
		}
		if stop {
			break
		}
	}

	fmt.Println("ooooooooooooooooooooooooooooooooooooooooooooo END-APP")
	fmt.Println("ooooooooooooooooooooooooooooooooooooooooooooo END-APP")
	fmt.Println("ooooooooooooooooooooooooooooooooooooooooooooo END-APP")

	if !stop {

		str := fmt.Sprintf("===== All steps are running")
		vutils.LogInfo(str)
		fmt.Println(str)

		select {
		case err = <-vomni.StepErr:
			str := fmt.Sprintf("Steps need to be closed due to an err - %q", err)
			vutils.LogErr(fmt.Errorf(str))
			stop = true

		case done = <-vomni.RootDone:
			stop = true
		}
	}

	l := ""

	switch done {
	case vomni.DoneShutdown:
		l = "SHUTDOWN"
	case vomni.DoneExit:
		l = "EXIT"
	case vomni.DoneRestart:
		l = "RESTART"
	case vomni.DoneReboot:
		l = "REBOOT"
	case vomni.DoneUpdateCode:
		l = "UPDATE CODE"
	default:
		l = "nezinu"
	}

	fmt.Println("Tagad jābeidz...", l)

	if (done > 0) && (vomni.DoneExit != done) {
		for count := vomni.StepCount(); count > 0; count-- {
			// let's do Post of each step starting from the last one
			ind := count - 1

			locDone := make(chan bool)
			thisS := steps[stepSequence[ind]]
			go thisS.StepPost(locDone)
			<-locDone

			str := fmt.Sprintf("===== Step %q -> closed", thisS.StepName())
			vutils.LogStr(vomni.LogInfo, str)
			fmt.Println(str)
		}
	}

	if stop {
		if nil != err {
			vomni.RootErr <- err
		}
		if 0 != done {
			chanDone <- done
		}
		return
	}
}
