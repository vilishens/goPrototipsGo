package pointrun

import (
	"fmt"
	"sort"
	vomni "vk/omnibus"
)

func AllPointData() (data vomni.WebAllPointData) {

	pts := make(map[string]vomni.WebPointData)

	list := []string{}

	for k, v := range Points {

		list = append(list, k)

		d := vomni.WebPointData{}

		d.Point = k
		d.State = v.Run[1].GetState()
		d.Type = v.Point.Type

		d.Signed = 0 != (v.Point.State & vomni.PointStateSigned)
		d.Disconnected = 0 != (v.Point.State & vomni.PointStateDisconnected)

		// jāizdomā, ko darīt ar punkta datiem
		state := v.Run[1].GetState()
		d.Frozen = 0 != (state & vomni.PointStateFrozen)

		//st := v.Run
		//keys := reflect.ValueOf(st).MapKeys()

		//state := 2

		d.CfgList = vomni.CfgListSequence
		d.CfgInfo = webCfgInfo(d.CfgList) // ziņas par konfigurācijām: pāšlaik konfigurāciju nosaukumi

		d.CfgDefault = make(map[int]interface{})
		d.CfgRun = make(map[int]interface{})
		d.CfgSaved = make(map[int]interface{})
		d.CfgState = make(map[int]interface{})
		d.CfgIndex = make(map[int]interface{})

		// get all configurations of this poist
		for _, cc := range d.CfgList {
			cfgDef, cfgRun, cfgSaved, cfgIndex, cfgState := Points[k].Run[cc].GetCfgs()

			d.CfgDefault[cc], d.CfgRun[cc], d.CfgSaved[cc], d.CfgIndex[cc], d.CfgState[cc] =
				cfgDef, cfgRun, cfgSaved, cfgIndex, cfgState //Points[k].Run[cc].GetCfgs()

			_, _, _ = cfgDef, cfgRun, cfgSaved
		}

		pts[k] = d
	}

	sort.Strings(list)

	data.List = list
	data.Data = pts

	return
}

func webCfgInfo(list []int) (d map[int]vomni.CfgPlusData) {

	d = make(map[int]vomni.CfgPlusData)

	for _, v := range list {
		dd := vomni.CfgPlusData{}

		dd.Name = vomni.PointCfgData[v].CfgStr

		d[v] = dd
	}

	return
}

func WebSent(todo int, point string, data interface{}) {

	fmt.Println("$$$\n$$$\n$$$ Jāizdomā, ko darīt ar konstantēm PointCmdBits (tā par lielu priekš Raspja) pagaidām noņēmu Raspja versijā vienu hexu\n$$$\n$$$\n$$$")

	cmd := todo & vomni.PointCmdBits
	switch cmd {
	case vomni.PointCmdLoadCfgIntoPoint, vomni.PointCmdSaveCfg, vomni.PointCmdFreezeOn, vomni.PointCmdFreezeOff:
		cfg := todo & vomni.PointCmdOptionBits
		Points[point].Run[cfg].ReceiveWeb(cmd, data)

	default:
		str := fmt.Sprintf("\n\nDon't know what to do with %08X for %s\n\n", todo, point)
		panic(str)
	}
}

func WebCmd(todo int, point string) {
	cmd := todo & vomni.PointCmdBits
	cfg := todo & vomni.PointCmdOptionBits
	Points[point].Run[cfg].Cmd(cmd)
}
