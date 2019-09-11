package cfg

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	vomni "vk/omnibus"
	vutils "vk/utils"
)

var Final CfgFinalData

func init() {

	Final.StationName = ""

	Final.LogMainPath = ""

	Final.PortUDPInternal = -1
	Final.PortSSHInternal = -1
	Final.PortWEBInternal = -1
	Final.PortSSHExternal = -1
	Final.PortWEBExternal = -1

	Final.RotateMainTmpl = ""
	Final.RotatePointDataTmpl = ""
	Final.RotatePointInfoTmpl = ""
	Final.RotateRunCfg = ""
	Final.RotateRunSecs = -1
	Final.RotateStatusFileName = ""

	Final.WebStaticPrefix = ""
	Final.WebStaticDir = ""
	Final.WebTemplateDir = ""

	Final.IPExternalAddressCmds = []string{}
	Final.NetExternalRequirement = -1

	Final.PointConfigDefaultFile = ""
	Final.PointConfigFile = ""

	Final.MessageEmailAddress = ""
	Final.MessageSMTPHost = ""
	Final.MessageSMTPUser = ""
	Final.MessageSMTPPass = ""
	Final.MessageSMTPPort = -1
}

func Cfg(chDone chan bool, chErr chan error) {

	locDone := make(chan bool)
	locErr := make(chan error)

	go load(locDone, locErr)

	select {
	case err := <-locErr:
		chErr <- err
	case <-locDone:
		chDone <- true
	}
}

func load(chDone chan bool, chErr chan error) {

	var err error

	if err = loadCfg(); nil != err {
		vutils.LogStr(vomni.LogErr, err.Error())
		chErr <- err
		return
	}

	chDone <- true
}

func loadCfg() (err error) {

	full := ""
	err = error(nil)
	if full, err = cfgPath(); nil != err {
		return
	}

	if "" == full {
		err = fmt.Errorf("There is no Application configuration")
		return vutils.ErrFuncLine(err)
	}

	data, err := readCfg(full)
	if nil != err {
		return
	}

	if err = data.Put(); nil != err {
		return
	}

	return
}

func cfgPath() (path string, err error) {
	// configuration data path found in CLI flags
	cpath := flag.Lookup(vomni.CliCfgPathFld).Value.String()

	if "" == cpath {
		return
	}

	path = vutils.FileAbsPath(cpath, "")

	ok := false
	if ok, err = vutils.PathExists(path); !ok {
		err = fmt.Errorf("File \"%s\" doesn't exist", path)
		err = vutils.ErrFuncLine(err)
	}

	return
}

func readCfg(full string) (data CfgData, err error) {

	data = CfgData{}
	if ok, err := vutils.PathExists(full); !ok {
		return data, vutils.ErrFuncLine(err)
	}

	raw, err := ioutil.ReadFile(full)
	if err != nil {
		return data, vutils.ErrFuncLine(err)
	}

	if err = json.Unmarshal(raw, &data); nil != err {
		return data, vutils.ErrFuncLine(err)
	}

	return data, err
}

func (c *CfgData) Put() (err error) {

	if (nil == err) && ("" != c.StationName) {
		Final.StationName = c.StationName
	}

	// hard coded Main log file path
	Final.LogMainPath = filepath.Join(vomni.RootPath, vomni.LogMainPath)

	// the point log base path
	if (nil == err) && ("" != c.LogPointPath) {
		Final.LogPointPath = c.LogPointPath
	}
	// rotation of logs
	if (nil == err) && ("" != c.RotateMainTmpl) {
		Final.RotateMainTmpl = c.RotateMainTmpl
	}
	if (nil == err) && ("" != c.RotatePointDataTmpl) {
		Final.RotatePointDataTmpl = c.RotatePointDataTmpl
	}
	if (nil == err) && ("" != c.RotatePointInfoTmpl) {
		Final.RotatePointInfoTmpl = c.RotatePointInfoTmpl
	}
	if (nil == err) && ("" != c.RotateRunCfg) {
		Final.RotateRunCfg = c.RotateRunCfg
	}
	if (nil == err) && ("" != c.RotateRunSecs) {
		Final.RotateRunSecs, err = strconv.Atoi(c.RotateRunSecs)
	}
	if (nil == err) && ("" != c.RotateStatusFileName) {
		Final.RotateStatusFileName = c.RotateStatusFileName
	}

	// internal ports
	if (nil == err) && ("" != c.PortUDPInternal) {
		Final.PortUDPInternal, err = strconv.Atoi(c.PortUDPInternal)
	}
	if (nil == err) && ("" != c.PortSSHInternal) {
		Final.PortSSHInternal, err = strconv.Atoi(c.PortSSHInternal)
	}
	if (nil == err) && ("" != c.PortWEBInternal) {
		Final.PortWEBInternal, err = strconv.Atoi(c.PortWEBInternal)
	}
	// external ports
	if (nil == err) && ("" != c.PortSSHExternal) {
		Final.PortSSHExternal, err = strconv.Atoi(c.PortSSHExternal)
	}
	if (nil == err) && ("" != c.PortWEBExternal) {
		Final.PortWEBExternal, err = strconv.Atoi(c.PortWEBExternal)
	}

	// WEB configuration
	if (nil == err) && ("" != c.WebStaticPrefix) {
		Final.WebStaticPrefix = c.WebStaticPrefix
	}
	if (nil == err) && ("" != c.WebStaticDir) {
		Final.WebStaticDir = c.WebStaticDir
	}

	if (nil == err) && ("" != c.WebTemplateDir) {
		Final.WebTemplateDir = c.WebTemplateDir
	}

	// External net settings
	if (nil == err) && (0 < len(c.IPExternalAddressCmds)) {
		Final.IPExternalAddressCmds = make([]string, len(c.IPExternalAddressCmds))
		copy(Final.IPExternalAddressCmds, c.IPExternalAddressCmds)
	}
	if (nil == err) && ("" != c.NetExternalRequirement) {
		Final.NetExternalRequirement, err = strconv.Atoi(c.NetExternalRequirement)
	}

	// Point configuration file
	if (nil == err) && ("" != c.PointConfigDefaultFile) {
		Final.PointConfigDefaultFile = c.PointConfigDefaultFile
	}
	if (nil == err) && ("" != c.PointConfigFile) {
		Final.PointConfigFile = c.PointConfigFile
	}

	if (nil == err) && ("" != c.MessageEmailAddress) {
		Final.MessageEmailAddress = c.MessageEmailAddress
	}
	if (nil == err) && ("" != c.MessageSMTPHost) {
		Final.MessageSMTPHost = c.MessageSMTPHost
	}
	if (nil == err) && ("" != c.MessageSMTPUser) {
		Final.MessageSMTPUser = c.MessageSMTPUser
	}
	if (nil == err) && ("" != c.MessageSMTPPass) {
		Final.MessageSMTPPass = c.MessageSMTPPass
	}
	if (nil == err) && ("" != c.MessageSMTPPort) {
		Final.MessageSMTPPort, err = strconv.Atoi(c.MessageSMTPPort)
	}
	return
}
