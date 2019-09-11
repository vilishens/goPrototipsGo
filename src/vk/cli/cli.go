package cli

import (
	"flag"
	vomni "vk/omnibus"
)

var CliCfgDefaultPath string

func init() {
	flag.StringVar(&CliCfgDefaultPath, vomni.CliCfgPathFld, vomni.CfgDefaultPath, "configuration file path")

	flag.Parse()
}

func Init() (err error) {
	return
}
