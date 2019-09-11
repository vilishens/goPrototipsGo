package web

import (
	"html/template"
	"path/filepath"

	vparams "vk/params"
)

var tmpls = new(template.Template)
var tmplFiles []string
var tmplFuncs template.FuncMap

var tmplPath string // the base path of templates

func setTmpls() {

	tmplPath = vparams.Params.WebTemplateDir

	initTmpls()
	initFuncs()

	tmpls = template.Must(template.New("app").Funcs(tmplFuncs).ParseFiles(tmplFiles...))
}

func initTmpls() {
	tmplFiles = []string{
		tmplPath + "/base/base-footer.tmpl",
		tmplPath + "/base/base-header.tmpl",
		tmplPath + "/base/base-js.tmpl",
		tmplPath + "/base/base-navigation.tmpl",
	}

	addPage(&tmplFiles, tmplPath, "home")
	addPage(&tmplFiles, tmplPath, "about")
	addPage(&tmplFiles, tmplPath, "points/pointlist")
	addPage(&tmplFiles, tmplPath, "points/pointcfg/relayinterval")

	/*
		addPage(&tmplFiles, tmplPath, "login")
		addPage(&tmplFiles, tmplPath, "points/pointlist")
		addPage(&tmplFiles, tmplPath, "points/pointcfg/relayonoffinterval")
	*/
}

func addPage(files *([]string), path string, page string) {

	dir := filepath.Dir(page)
	base := filepath.Base(page)

	newF := filepath.Join(path, dir, base, base)
	*files = append(*files, newF+".tmpl")
	*files = append(*files, newF+"-body.tmpl")
}

func initFuncs() {
	tmplFuncs = make(map[string]interface{})
	tmplFuncs["stationName"] = stationName
	tmplFuncs["allPointData"] = allPointData
	tmplFuncs["pointData"] = pointData
	//	tmplFuncs["increment1"] = increment1
	//bija	tmplFuncs["pointCfgJsFile"] = pointCfgJsFile
	tmplFuncs["webPrefix"] = webPrefix
	//	tmplFuncs["webPref"] = webPref
}
