package pointconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	vomni "vk/omnibus"
	vparams "vk/params"
	vutils "vk/utils"
)

// `Šitie 4 nav it kā nav najadzīgi`
var PointsAllJSON JSONData
var PointsAllDefaultJSON JSONData
var PointsAllData AllPointCfgData
var PointsAllDefaultData AllPointCfgData

var AllPointData AllCfgData

func init() {
	//	PointsAllData = make(map[string]PointCfgData)
	//	PointsAllDefaultData = make(map[string]PointCfgData)
	//	PointsAllJSON = CfgJSONData{}
	//	PointsAllDefaultJSON = CfgJSONData{}

	AllPointData = AllCfgData{}
	AllPointData.Default = make(map[string]CfgPointData)
	AllPointData.DefaultJSON = CfgFileJSON{}
	AllPointData.Running = make(map[string]CfgPointData)
	AllPointData.RunningJSON = CfgFileJSON{}
}

func loadAllCfgFiles() (err error) {

	// check if the running configuration file exist, if not - clone the default configuration file
	if err = verifyCfgFiles(); nil != err {
		err = vutils.ErrFuncLine(fmt.Errorf("Couldn't verify existance of configuration files - %v", err))
		return
	}

	path := vutils.FileAbsPath(vparams.Params.PointConfigFile, "")

	fmt.Printf("\n\n\n******* RUNNING PATH %q\n\n\n", path)

	if AllPointData.Running, AllPointData.RunningJSON, err = loadCfgFile(path); nil != err {
		err = vutils.ErrFuncLine(fmt.Errorf("Active point configuration load failed - %v", err))
		return
	}

	path = vutils.FileAbsPath(vparams.Params.PointConfigDefaultFile, "")

	fmt.Printf("\n\n\n******* DEFAULT PATH %q\n\nDATA%+v\nJSON%+v\n", path, AllPointData.Default, AllPointData.DefaultJSON)

	if AllPointData.Default, AllPointData.DefaultJSON, err = loadCfgFile(path); nil != err {
		err = vutils.ErrFuncLine(fmt.Errorf("Default point configuration load failed - %v", err))
		return
	}

	return
}

func loadCfgFile(path string) (data CfgFileData, json CfgFileJSON, err error) {

	fmt.Printf("\n\n\n******* PAth 4 JSON %q\n\n", path)

	// read the configuration file provided by the path
	if json, err = getCfgJSON(path); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	fmt.Printf("=========================================== JSON =====================================\n%+v\n", json)

	// put the file data into cofiguration structures
	if data, err = json.putCfgJSON4Run(); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	fmt.Printf("=========================================== DATA =====================================\n%+v\n", json)

	fmt.Println("Marketa Davidova @@@@@@@@@@@@@\n", json)

	return
}

func getCfgJSON(path string) (data CfgFileJSON, err error) {

	if err = vutils.ReadJson(path, &data); nil != err {
		return CfgFileJSON{}, vutils.ErrFuncLine(err)
	}

	return
}

func (d CfgFileJSON) putCfgJSON4Run() (data CfgFileData, err error) {

	data = make(map[string]CfgPointData)

	for k, v := range d {

		newStorage := CfgPointData{}

		// add RelayInterval configuration
		if newStorage, err = v.putRelayIntervalJSON4Run(data[k]); nil != err {
			err = vutils.ErrFuncLine(fmt.Errorf("Relay Interval configuration Error - %s", err.Error()))
			return
		}
		data[k] = newStorage

		// add all Temperature Relay configurations for Run
		if newStorage, err = v.putTempRelayJSON4Run(data[k]); nil != err {
			err = vutils.ErrFuncLine(fmt.Errorf("Temperature Relay configuration Error - %s", err.Error()))
			return
		}
		data[k] = newStorage

	}

	return
}

func verifyCfgFiles() (err error) {
	if has, _ := vutils.PathExists(vparams.Params.PointConfigFile); !has {
		err = vutils.FileCopy(vparams.Params.PointConfigDefaultFile, vparams.Params.PointConfigFile)
	}
	return
}

func loadPointCfg() (data JSONData, err error) {

	if has, _ := vutils.PathExists(vparams.Params.PointConfigFile); !has {
		if err := vutils.FileCopy(vparams.Params.PointConfigDefaultFile, vparams.Params.PointConfigFile); nil != err {
			return JSONData{}, vutils.ErrFuncLine(err)
		}
	}

	if err = vutils.ReadJson(vparams.Params.PointConfigFile, &data); nil != err {
		return JSONData{}, vutils.ErrFuncLine(err)
	}

	return
}

func loadPointDefaultCfg() (data JSONData, err error) {

	if err = vutils.ReadJson(vparams.Params.PointConfigDefaultFile, &data); nil != err {
		return JSONData{}, vutils.ErrFuncLine(err)
	}

	return
}

//##############################################################################
//##############################################################################
//##############################################################################
//##############################################################################
//##############################################################################
//##############################################################################

func GetPointCfg(chGoOn chan bool, chDone chan int, chErr chan error) {
	locDone := make(chan bool)
	locErr := make(chan error)

	go preparePointCfg(locDone, locErr)

	select {
	case err := <-locErr:
		vutils.LogErr(err)
		chErr <- vutils.ErrFuncLine(err)
	case <-locDone:
		chGoOn <- true
	}

	fmt.Println(PointsAllData.String())
}

func preparePointCfg(doneCh chan bool, errCh chan error) {

	var err error

	if err = loadAllCfgFiles(); nil != err {
		err = vutils.ErrFuncLine(err)
		errCh <- err
		return
	}

	fmt.Printf("===================\n===================\n===================\nDEFAULT\n%+v\nRUNNING%+v\n===================\n===================\n",
		AllPointData.Default, AllPointData.Running)

	doneCh <- true
	return

	//##############################################################################
	//##############################################################################
	//##############################################################################

	if PointsAllJSON, err = loadPointCfg(); nil == err {
		err = PointsAllJSON.putCfg4Run()
	}

	if nil != err {
		errCh <- err
		return
	}

	if PointsAllDefaultJSON, err = loadPointDefaultCfg(); nil == err {
		err = PointsAllDefaultJSON.putCfgDefault4Run()
	}

	if nil != err {
		errCh <- err
		return
	}

	doneCh <- true
}

func (d JSONData) putCfg4Run() (err error) {

	for k, v := range d {
		if _, has := PointsAllData[k]; !has {
			PointsAllData[k] = CfgPointData{}
		}

		var newStruct CfgPointData
		if v.RelIntervalJSON.hasCfgRelInterval() {
			if newStruct, err = v.RelIntervalJSON.putCfg4Run(PointsAllData[k]); nil != err {
				err = vutils.ErrFuncLine(fmt.Errorf("Relay Interval configuration Error - %s", err.Error()))
				return
			} else {
				PointsAllData[k] = newStruct
			}
		}
	}

	return
}

func (d JSONData) putCfgDefault4Run() (err error) {

	for k, v := range d {
		if _, has := PointsAllDefaultData[k]; !has {
			PointsAllDefaultData[k] = CfgPointData{}
		}

		var newStruct CfgPointData
		if v.RelIntervalJSON.hasCfgRelInterval() {
			if newStruct, err = v.RelIntervalJSON.putCfg4Run(PointsAllDefaultData[k]); nil != err {
				err = vutils.ErrFuncLine(fmt.Errorf("Relay Interval Default configuration Error - %s", err.Error()))
				return
			} else {
				PointsAllDefaultData[k] = newStruct
			}
		}
	}

	return
}

func (d JSONRelIntervalStruct) hasCfgRelInterval() (has bool) {
	if 0 < len(d.Start) {
		return true
	}

	if 0 < len(d.Base) {
		return true
	}

	if 0 < len(d.Finish) {
		return true
	}

	return false
}

func (d CfgFileJSON) Save() (err error) {

	data, _ := json.Marshal(d)
	path := vutils.FileAbsPath(vparams.Params.PointConfigFile, "")
	if err = ioutil.WriteFile(path, data, vomni.FilePermissions); nil != err {
		return vutils.ErrFuncLine(err)
	}

	return
}
