package github

import (
	"fmt"
	"runtime"
)

func GetArchiveName() string {
	return fmt.Sprintf("%s-%s.zip", runtime.GOOS, runtime.GOARCH)

}

func GetReleaseUrl() string {
	return fmt.Sprintf("https://github.com/alshdavid/alshx/releases/latest/download/%s", GetArchiveName())
}
