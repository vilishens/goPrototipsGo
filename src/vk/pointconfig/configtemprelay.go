package pointconfig

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	vomni "vk/omnibus"
	vutils "vk/utils"
)

func (d JSONTempRelay) hasCfgTempInterval() (has bool) {
	if 0 < len(d.Conditions) {
		return true
	}

	return false
}

func (d JSONTempRelay) addCfgTempInterval2Run(dst CfgPointData) (newStorage CfgPointData, err error) {
	if newStorage, err = d.putCfg4Run(dst); nil != err {
		err = vutils.ErrFuncLine(fmt.Errorf("Temperature Relay configuration transformation to Run Error - %s", err.Error()))
		return
	}

	return
}

func (d JSONTempRelay) putCfg4Run(dst CfgPointData) (newDst CfgPointData, err error) {

	newDst = CfgPointData{}
	newD := RunTempRelay{}
	var fl float64

	if newD.Conditions, err = d.Conditions.putConditions4Run(); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	if fl, err = strconv.ParseFloat(d.Delta, 32); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}
	newD.Delta = float32(math.Round(fl*100) / 100)

	newD.Fahrenheit = false
	if d.Fahrenheit != "" && "1" == strings.Trim(d.Fahrenheit, " ") {
		newD.Fahrenheit = true
	}

	if newD.Seconds, err = vutils.ConfInterval2Seconds(d.Interval); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	newD.Handler = strings.Trim(d.Handler, " ")

	if newD.Gpio, err = strconv.Atoi(d.Gpio); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	if newD.State, err = strconv.Atoi(d.State); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	if newD.Start, err = strconv.Atoi(d.Start); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	if newD.Finish, err = strconv.Atoi(d.Finish); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	newDst = dst

	newDst.List |= vomni.CfgTypeTempRelay
	newDst.Cfg.TempRelay = append(newDst.Cfg.TempRelay, newD)
	newDst.CfgSaved.TempRelay = append(newDst.CfgSaved.TempRelay, newD)

	return
}

func (d JSONConditions) putConditions4Run() (newD RunConditions, err error) {

	newD = RunConditions{}

	for _, v := range d {
		newC := RunCondition{}

		var fl float64

		if fl, err = strconv.ParseFloat(v.MinTemp, 32); nil != err {
			err = vutils.ErrFuncLine(err)
			return
		}
		newC.MinTemp = float32(math.Round(fl*100) / 100)

		if fl, err = strconv.ParseFloat(v.MaxTemp, 32); nil != err {
			err = vutils.ErrFuncLine(err)
			return
		}
		newC.MaxTemp = float32(math.Round(fl*100) / 100)

		if newC.Mask, err = strconv.Atoi(v.Mask); nil != err {
			err = vutils.ErrFuncLine(err)
			return
		}

		newD = append(newD, newC)
	}

	return
}

func (d JSONPointData) putTempRelayJSON4Run(storage CfgPointData) (newStorage CfgPointData, err error) {

	// add TempRelay (separate) configuration
	if d.TempRelayJSON.hasCfgTempInterval() {
		if newStorage, err = d.TempRelayJSON.addCfgTempInterval2Run(storage); nil != err {
			err = vutils.ErrFuncLine(fmt.Errorf("Couldn't prepare Temperature Relay configuration for Run - %s", err.Error()))
			return
		}

		storage = newStorage
	}

	// add TempRelay (array) configurations
	for _, v := range d.TempRelayArrayJSON {
		if v.hasCfgTempInterval() {
			if newStorage, err = v.addCfgTempInterval2Run(storage); nil != err {
				err = vutils.ErrFuncLine(fmt.Errorf("Couldn't prepare Temperature Relay array configurations for Run - %s", err.Error()))
				return
			}

			storage = newStorage
		}
	}

	return storage, err
}
