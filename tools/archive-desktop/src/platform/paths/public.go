package paths

import (
	"os"
	"path/filepath"
)

func GetPublicDesktop() string {
	profilesPath := os.Getenv("PUBLIC")
	if profilesPath == "" {
		return ""
	}
	return filepath.Join(profilesPath, "Desktop")
}
