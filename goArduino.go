package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	vomni "vk/omnibus"
	sall "vk/steps/allsteps"
	vutils "vk/utils"
)

func init() {
	root()
}

func main() {

	end := false
	endCd := -1

	for !end {
		if 0 > endCd {
			vutils.LogStr(vomni.LogInfo, "***** App - START *****")
		}

		endCd := runApp()

		switch endCd {
		case vomni.DoneExit:
			vutils.LogStr(vomni.LogInfo, "***** App - EXIT *****")
			end = true
		case vomni.DoneRestart:
			vutils.LogStr(vomni.LogInfo, "***** App - RESTART *****")
			end = true
		case vomni.DoneStop:
			vutils.LogStr(vomni.LogInfo, "***** App - STOP *****")
			end = true
		case vomni.DoneShutdown:
			vutils.LogStr(vomni.LogInfo, "***** App - SHUTDOWN *****")
			end = true
		case vomni.DoneError:
			vutils.LogStr(vomni.LogInfo, "***** App - ERROR *****")
			end = true
		case vomni.DoneReboot:
			vutils.LogStr(vomni.LogInfo, "***** App - REBOOT *****")
			end = true
		case vomni.DoneUpdateGo:
			vutils.LogStr(vomni.LogInfo, "***** App - UPDATE GO CODE *****")

			fmt.Println("###\n###\n###\n###\n MARIO ", endCd)

			end = true
		default:
			str := fmt.Sprintf("***** App - unknown Exit code %d *****", endCd)
			vutils.LogStr(vomni.LogInfo, str)
			fmt.Println(str)
		}

		if end {

			fmt.Println("###\n###\n###\n###\n EXIT code ", endCd)

			os.Exit(endCd)
		}
	}
}

func runApp() (cd int) {
	chDone := make(chan int)

	go sall.DoSteps(chDone)

	select {
	case err := <-vomni.RootErr:
		fmt.Printf("App finished due to an error ---> %v\n", err)
		cd = vomni.DoneError
		break
	case cd = <-chDone:
		str := fmt.Sprintf("***** App - received code %d *****", cd)
		vutils.LogStr(vomni.LogInfo, str)
		fmt.Println(str)
		break
	}

	return
}

func root() {
	rootPath()
	rootLog()
}

func rootPath() {
	// It is necessary to keep the root caller path to create
	// correct file paths further
	if _, rootFile, _, ok := runtime.Caller(0); !ok {
		err := fmt.Errorf("Could not get Root Path")
		log.Fatal(err)
	} else {
		vomni.RootPath = filepath.Dir(rootFile)
	}

	return
}

func rootLog() {
	var err error

	path := filepath.Join(vomni.RootPath, vomni.LogMainPath)

	vomni.LogMainFile, err = vutils.OpenFile(path, vomni.LogFileFlags, vomni.LogUserPerms)
	if nil != err {
		log.Fatal(fmt.Errorf("Could not open the main log file --- %v", err))
	}

	vomni.LogData = vutils.LogNew(vomni.LogMainFile, vomni.LogPrefixData)
	vomni.LogErr = vutils.LogNew(vomni.LogMainFile, vomni.LogPrefixErr)
	vomni.LogFatal = vutils.LogNew(vomni.LogMainFile, vomni.LogPrefixFatal)
	vomni.LogInfo = vutils.LogNew(vomni.LogMainFile, vomni.LogPrefixInfo)
}
