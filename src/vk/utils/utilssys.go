package utils

import (
	"os/exec"
)

func DoCmd(cmd string, chStr chan string, chErr chan error) {
	res, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		chErr <- err
	} else {
		chStr <- string(res)
	}
}
