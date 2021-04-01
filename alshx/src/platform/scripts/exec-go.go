package scripts

import (
	"alshx/src/platform/logging"
	"os"
	"os/exec"
	"path/filepath"
)

func execGo(
	logger *logging.Logger,
	cmd *exec.Cmd,
	config *Meta,
	args []string,
	folderPath string,
) *exec.Cmd {
	if !hasCommand("go") {
		logger.Log("Go is not installed")
		os.Exit(1)
	}
	cmdPath := []string{"go", "run", filepath.Join(folderPath, config.Entrypoint)}
	cmdPath = append(cmdPath, config.Args...)
	cmdPath = append(cmdPath, args...)
	logger.Info("Command:", cmdPath)
	cmd = exec.Command(cmdPath[0], cmdPath[1:]...)
	return cmd
}
