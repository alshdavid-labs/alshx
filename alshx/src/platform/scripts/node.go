package scripts

import (
	"alshx/src/platform/files"
	"alshx/src/platform/logging"
	"os"
	"os/exec"
	"path/filepath"
)

func execNode() {

}

func installNodeModules(
	logger *logging.Logger,
	folderPath string,
) {
	if files.NotExists(filepath.Join(folderPath, "node_modules")) {
		logger.Log("Installing Node Modules")
		cmd := exec.Command("yarn")
		cmd.Dir = folderPath
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
