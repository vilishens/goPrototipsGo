package rotate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"time"
	vomni "vk/omnibus"
	vparams "vk/params"
	vutils "vk/utils"
)

var myLoggers []ActiveLog

func init() {
	myLoggers = []ActiveLog{}
}

func PointLogger(logFil string, tmplFile string) {
	//	myLoggers = append(myLoggers, ActiveLog{Path: logPath, File: f, Loggers: logs})
}

func addMainLogs() {

	mainList := []*log.Logger{vomni.LogData, vomni.LogErr, vomni.LogFatal, vomni.LogInfo}
	myLoggers = append(myLoggers, ActiveLog{Path: vparams.Params.LogMainPath, File: vomni.LogMainFile, Loggers: mainList})
}

func MainStart(chGoOn chan bool, chDone chan int, chErr chan error) {

	fmt.Println("!!!!!!!!!!!!! ROTATION NAV Vēl līdz galam izbaudīta !!!!!!!!!!!")

	// Log Main loggers have created already it's necessary to put its in the list of loggers
	addMainLogs()

	// start the brand new rotation configuration with the Main Log file
	err := SetRotateCfg(vparams.Params.LogMainPath, vparams.Params.RotateMainTmpl, vparams.Params.RotateRunCfg, true)
	if nil != err {
		err = vutils.ErrFuncLine(err)
		vutils.LogStr(vomni.LogErr, err.Error())
		chErr <- err
		return
	}

	chGoOn <- true

	locDone := make(chan int)
	locErr := make(chan error)
	go runRotate(locDone, locErr)

	select {
	case err = <-locErr:
		err = vutils.ErrFuncLine(err)
		vutils.LogStr(vomni.LogErr, err.Error())
		chErr <- err
		return
	case <-locDone:
		chDone <- vomni.DoneStop
		return
	}
}

// 	file2Rotate - file for log records
//	cfgTmpl2Use - the rotate configuration template file
//  cfg2RunFile - the rotation configuration data file (all separate logger configuration in one place - set in app.cfg)
func SetRotateCfg(file2Rotate string, cfgTmpl2Use string, cfg2RunFile string, newRotation bool) (err error) {

	usr := new(user.User)
	/*
	 * find the current user data (user name and user group)
	 * to create this particular rotation configuration
	 */
	usr, err = user.Current()
	if nil != err {
		return vutils.ErrFuncLine(err)
	}

	name := usr.Username

	group := ""
	if usrGrp, err := user.LookupGroupId(usr.Gid); nil != err {
		return vutils.ErrFuncLine(err)
	} else {
		group = usrGrp.Name
	}

	/*
	 * Prepare rotation configuration
	 * from the configuration template
	 */
	var format []byte
	// read necessary data file rotation configuration template
	if format, err = ioutil.ReadFile(cfgTmpl2Use); nil != err {
		return vutils.ErrFuncLine(err)
	}

	// put required data (user name and user group) into configuration template
	str := fmt.Sprintf(string(format), file2Rotate, name, group)

	/*
	 * delete the existing running configuration
	 * if it is required to start with the brand new configuration file
	 */
	if newRotation {
		has, err := vutils.PathExists(cfg2RunFile)
		if nil != err {
			return vutils.ErrFuncLine(err)
		}
		if has {
			if err = vutils.FileDelete(cfg2RunFile); nil != err {
				return vutils.ErrFuncLine(err)
			}
		}
	}

	/*
	 * use configuration prepared from the template
	 * to put it in the running configuration
	 * or to add (in case the rotation was started - newRotation is false)
	 */
	if err = vutils.FileAppend(cfg2RunFile, str); nil != err {
		return vutils.ErrFuncLine(err)
	}

	return
}

func runRotate(chDone chan int, chErr chan error) {

	if 0 >= vparams.Params.RotateRunSecs {
		chErr <- vutils.ErrFuncLine(fmt.Errorf("\nWrong point rotation interval %d", vparams.Params.RotateRunSecs))
		return
	}

	for {
		timer := time.NewTimer(time.Duration(vparams.Params.RotateRunSecs) * time.Second)

		// rotate
		if err := rotateFiles(); nil != err {
			chErr <- vutils.ErrFuncLine(fmt.Errorf("\nRotation command failure -- %s", err.Error()))
			return
		}

		select {
		case <-timer.C:

			vutils.LogStr(vomni.LogInfo, "Time to rotate log files")

			timeStr := time.Now().Format("2006-01-02 15:04:05 -07:00 MST")
			str := fmt.Sprintf("==>>>>> %s <<<<<< ROTATE", timeStr)
			fmt.Println(str)
		}
	}
}

func rotateFiles() (err error) {
	// rotate files if necessary
	if err = runRotateCmd(); nil != err {
		err = vutils.ErrFuncLine(fmt.Errorf("\nRotation command failure -- %s", err.Error()))
		vutils.LogStr(vomni.LogErr, err.Error())
		return
	}

	for k, v := range myLoggers {
		// set the log file to the original path
		// as it is renamed in case of rotation
		if myLoggers[k].File, err = LogReassignFile(v.File, v.Path); nil != err {
			err = vutils.ErrFuncLine(fmt.Errorf("Log file reassign failure -- %s", err.Error()))
			vutils.LogStr(vomni.LogErr, err.Error())
			return
		}

		// link loggers to the log file object
		for _, v1 := range v.Loggers {
			v1.SetOutput(myLoggers[k].File)
		}
	}

	return
}

func LogReassignFile(f *os.File, path string) (fNew *os.File, err error) {
	var perms os.FileMode
	flags := vomni.LogFileFlags

	if nil != f {
		var stat os.FileInfo

		stat, err = f.Stat()
		if nil != err {
			return nil, vutils.ErrFuncLine(fmt.Errorf("Could not get stat of the file %s", path))
		}

		perms = stat.Mode()

		if err = f.Close(); nil != err {
			return nil, vutils.ErrFuncLine(fmt.Errorf("Could not close the file %s", path))
		}
	} else {
		perms = vomni.LogUserPerms
	}

	return vutils.OpenFile(path, flags, perms)
}

func runRotateCmd() (err error) {
	//	find the local status file
	dirpath := filepath.Dir(vparams.Params.RotateRunCfg)
	statusF := filepath.Join(dirpath, vparams.Params.RotateStatusFileName)

	if has, _ := vutils.PathExists(statusF); has {
		err = vutils.FileDelete(statusF)
		if nil != err {
			return vutils.ErrFuncLine(err)
		}
	}

	// logrotate <conf.file> -s <localstatus.file>
	cmd := exec.Command("logrotate", vparams.Params.RotateRunCfg, "-s", statusF)

	if err = cmd.Run(); nil != err {
		return vutils.ErrFuncLine(err)
	}

	return
}

func RotatePointFilePath(key int, path string, point string, cfg int) (fPath string) {

	j := 0
	ending := ""
	for j <= key {

		if 0 == j {
			j = 1
		} else {
			j <<= 1
		}

		if 0 == key&j {
			continue
		}

		ending += "." + vomni.PointLogData[j].FileEnd
	}

	// rotate log data file
	return vutils.FileAbsPath(filepath.Join(path, point), vomni.PointCfgData[cfg].CfgStr+ending)

}

func RotatePointLoggers(key int) (loggers map[int]vomni.PointLogger) {

	j := 0
	loggers = make(map[int]vomni.PointLogger)
	for j <= key {
		if 0 == j {
			j = 1
		} else {
			j <<= 1
		}

		if 0 == key&j {
			continue
		}

		loggers[j] = vomni.PointLogger{LogPrefix: vomni.PointLogData[j].LogPrefix, Logger: nil}
	}

	return
}

func StartPointLoggers(point string, logs []vomni.PointLog) (err error) {

	for _, v := range logs {
		// add the log data file to the running configuration file
		if err = SetRotateCfg(v.LogFile, v.LogTmpl, vparams.Params.RotateRunCfg, false); nil != err {
			return vutils.ErrFuncLine(err)
		}

		list := []*log.Logger{}
		for _, v1 := range v.Loggers {
			list = append(list, v1.Logger)
		}

		myLoggers = append(myLoggers, ActiveLog{Path: v.LogFile, File: v.LogFilePtr, Loggers: list})
	}

	return
}