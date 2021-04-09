package paths

import (
	"os"
	"path/filepath"
)

func Executable() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
