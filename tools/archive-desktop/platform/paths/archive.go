package paths

import (
	"os"
	"path/filepath"
)

func GetArchivePath() (archivePath string) {
	archivePath = os.Getenv("ALSHX_DESKTOP_ARCHIVE_PATH")
	if archivePath != "" {
		return archivePath
	}
	homePath := GetHomePath()
	archivePath = filepath.Join(homePath, "Archive")
	return archivePath
}
