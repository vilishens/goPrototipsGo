package pointconfig

import (
	"fmt"
	"strconv"
	vomni "vk/omnibus"
	vutils "vk/utils"
)

func (d JSONRelIntervalStruct) putCfg4RunX(point string) (err error) {

	newD := RunRelIntervalStruct{}
	if newD.Start, err = d.Start.putCfg4Run(); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	if newD.Base, err = d.Base.putCfg4Run(); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	if newD.Finish, err = d.Finish.putCfg4Run(); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	tmpD := PointsAllData[point]
	tmpD.List |= vomni.CfgTypeRelayInterval
	tmpD.Cfg.RelInterv = newD
	tmpD.CfgSaved.RelInterv = newD
	PointsAllData[point] = tmpD

	return
}

func (d JSONRelIntervalStruct) newRelIntervalStruct() (newD RunRelIntervalStruct, err error) {
	newD = RunRelIntervalStruct{}
	if newD.Start, err = d.Start.putCfg4Run(); nil != err {
		err = vutils.ErrFuncLine(err)
		return RunRelIntervalStruct{}, err
	}

	if newD.Base, err = d.Base.putCfg4Run(); nil != err {
		err = vutils.ErrFuncLine(err)
		return RunRelIntervalStruct{}, err
	}

	if newD.Finish, err = d.Finish.putCfg4Run(); nil != err {
		err = vutils.ErrFuncLine(err)
		return RunRelIntervalStruct{}, err
	}

	return
}

func (d JSONRelIntervalStruct) putCfgDefault4Run(dst CfgPointData) (newDst CfgPointData, err error) {

	newDst = CfgPointData{}
	newD := RunRelIntervalStruct{}

	if newD, err = d.newRelIntervalStruct(); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	newDst = dst

	newDst.List |= vomni.CfgTypeRelayInterval
	newDst.Cfg.RelInterv = newD
	newDst.CfgSaved.RelInterv = newD

	return
}

func (d JSONRelIntervalStruct) putCfg4Run(dst CfgPointData) (newDst CfgPointData, err error) {

	newDst = CfgPointData{}
	newD := RunRelIntervalStruct{}

	if newD, err = d.newRelIntervalStruct(); nil != err {
		err = vutils.ErrFuncLine(err)
		return
	}

	newDst = dst

	newDst.List |= vomni.CfgTypeRelayInterval
	newDst.Cfg.RelInterv = newD
	newDst.CfgSaved.RelInterv = newD

	return
}

func (d JSONRelIntervalArray) putCfg4Run() (newD []RunRelInterval, err error) {

	newD = []RunRelInterval{}

	for _, v := range d {
		tmpD := RunRelInterval{Gpio: -1, State: -1, Seconds: 0}

		if "" != v.Gpio {
			if tmpD.Gpio, err = strconv.Atoi(v.Gpio); nil != err {
				err = vutils.ErrFuncLine(err)
				return
			}
		}

		if "" != v.State {
			if tmpD.State, err = strconv.Atoi(v.State); nil != err {
				err = vutils.ErrFuncLine(err)
				return
			}
		}

		if "" != v.Interval {
			if tmpD.Seconds, err = vutils.ConfInterval2Seconds(v.Interval); nil != err {
				err = vutils.ErrFuncLine(err)
				return
			}
		}

		newD = append(newD, tmpD)
	}

	return
}

func (d JSONPointData) putRelayIntervalJSON4Run(storage CfgPointData) (newStorage CfgPointData, err error) {
	// add Relay Interval (separate) configuration
	if d.RelIntervalJSON.hasCfgRelInterval() {
		if newStorage, err = d.RelIntervalJSON.putCfg4Run(storage); nil != err {
			err = vutils.ErrFuncLine(fmt.Errorf("Couldn't prepare Relay Interval configuration Error - %s", err.Error()))
			return
		}

		storage = newStorage
	}

	return storage, err
}
