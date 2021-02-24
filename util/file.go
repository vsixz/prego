package util

import (
	"github.com/vsixz/prego/cons"
	"os"
	"runtime"
)

func IsPathExist(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func HomeDir() string {
	res := os.Getenv("HOME")
	if runtime.GOOS == cons.WindowsOS {
		if len(res) > 0 {
			return res
		}

		if res = os.Getenv("USERPROFILE"); len(res) > 0 {
			return res
		}

		if homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"); len(homeDrive) > 0 && len(homePath) > 0 {
			res = homeDrive + homePath
			return res
		}
	}
	return res
}
