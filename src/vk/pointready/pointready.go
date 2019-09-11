package pointready

import (
	vomni "vk/omnibus"
)

func Prepare(chGoOn chan bool, chDone chan int, chErr chan error) {

	relayInterval()

	chGoOn <- true
	//	for {
	//		time.Sleep(vomni.DelayStepExec)
	//	}
}

func pointLoggers(point string, pType int) (loggers vomni.PointLog, err error) {
	/*

		loggers = make(map[int]*log.Logger)

		file := vomni.LogPointInfo[pType].File
		for _, v := range vomni.LogPointInfo[pType].List {

			fName := file + "." + v

			full := vutils.FileAbsPath(filepath.Join(vparams.Params.LogPointPath, point), fName)

			//		fmt.Printf("###\n###\n%s\n###\n###\n", full)

			var lg *log.Logger
			lg, err = vutils.LogNewPath(full, v+" ")
			if nil == err {
				loggers[v] = lg
			} else {
				err = vutils.ErrFuncLine(fmt.Errorf("Could not create Log %q of the point %q", v, point))
				vutils.LogErr(err)

				return
			}
		}
	*/
	return
}
