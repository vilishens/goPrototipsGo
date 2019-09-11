package utils

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJson(path string, dst interface{}) (err error) {

	if ok, err := PathExists(path); !ok {
		return ErrFuncLine(err)
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return ErrFuncLine(err)
	}

	if err = json.Unmarshal(raw, &dst); nil != err {
		return ErrFuncLine(err)
	}

	return
}
