package files

import "os"

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func NotExists(path string) bool {
	return !Exists(path)
}
