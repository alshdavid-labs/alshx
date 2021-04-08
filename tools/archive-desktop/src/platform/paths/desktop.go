package paths

import (
	"os"
	"path/filepath"
)

func GetDesktopPath() (desktopPath string) {
	desktopPath = os.Getenv("ALSHX_DESKTOP_PATH")
	if desktopPath != "" {
		return desktopPath
	}
	homePath := GetHomePath()
	desktopPath = filepath.Join(homePath, "Desktop")
	return desktopPath
}
