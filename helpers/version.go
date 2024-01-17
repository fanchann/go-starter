package helpers

import (
	"os/exec"
	"regexp"
)

func GetGoVersion() string {
	execRes, err := exec.Command("go", "version").Output()
	ErrorWithLog(err)

	getVer := regexp.MustCompile(`\d+\.\d+`)
	version := getVer.FindString(string(execRes))

	return version
}
