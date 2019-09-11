package utils

import (
	"fmt"
	"runtime"
)

func ErrFuncLine(err error) (newErr error) {

	if nil == err {
		return
	}

	pc, fn, line, _ := runtime.Caller(1)

	newErr = fmt.Errorf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	return
}
