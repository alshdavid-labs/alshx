package paths

import (
	"os"
	"runtime"
)

var homePath string

func GetHomePath() string {
	if homePath != "" {
		return homePath
	} else if runtime.GOOS == "windows" {
		homePath = os.Getenv("USERPROFILE")
	} else {
		homePath = os.Getenv("HOME")
	}
	return homePath
}
