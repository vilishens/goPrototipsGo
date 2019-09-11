package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	vomni "vk/omnibus"
)

func OpenFile(path string, fileFlags int, userPerms os.FileMode) (f *os.File, err error) {

	if err = FileDir(path); nil != err {
		return
	}

	f, err = os.OpenFile(path, fileFlags, userPerms)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}

	return
}

func FileDir(full string) (err error) {

	permDir := os.FileMode(vomni.DirPermissions)

	dirpath := filepath.Dir(full)

	if err = os.MkdirAll(dirpath, permDir); nil != err {
		return
	}

	return
}

func FileAbsPath(fPath string, file string) (full string) {

	abs := ""

	if !filepath.IsAbs(fPath) {
		abs = vomni.RootPath
	}

	abs = filepath.Join(abs, fPath, file)
	full = filepath.Clean(abs)

	return
}

func PathExists(full string) (exists bool, err error) {

	pathStr := FileAbsPath(full, "")

	if _, err = os.Stat(pathStr); os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

func FileDelete(full string) (err error) {
	return os.Remove(full)
}

func FileAppend(fullPath string, strAdd string) (err error) {

	permDir := os.FileMode(vomni.DirPermissions)
	permFile := os.FileMode(vomni.FileNonExecPermissions)

	dirpath := filepath.Dir(fullPath)

	if err = os.MkdirAll(dirpath, permDir); nil != err {
		return
	}

	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, permFile)
	if err != nil {
		return
	}

	defer f.Close()

	_, err = f.WriteString(strAdd)

	return
}

func FileCopy(src string, dst string) (err error) {

	var srcF, dstF *os.File

	fullSrc := FileAbsPath(src, "")

	if srcF, err = os.Open(fullSrc); nil != err {
		return
	}
	defer srcF.Close()

	fullDst := FileAbsPath(dst, "")

	dirOnly := filepath.Dir(fullDst)

	if err = os.MkdirAll(dirOnly, vomni.DirPermissions); nil != err {
		return
	}

	if dstF, err = os.Create(fullDst); nil != err { // creates if file doesn't exist
		return
	}
	defer dstF.Close()

	if _, err = io.Copy(dstF, srcF); nil != err { // check first var for number of bytes copied
		return
	}

	return dstF.Sync()
}
