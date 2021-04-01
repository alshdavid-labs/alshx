package scripts

import (
	"alshx/src/platform/logging"
	"os/exec"
)

func execGo(
	logger *logging.Logger,
	cmd *exec.Cmd,
	config *Meta,
	args []string,
) {
	if !hasCommand("go") {
		logger.Log("Go is not installed")
		return
	}
	cmdPath := []string{"go", "run", config.Entrypoint}
	cmdPath = append(cmdPath, config.Args...)
	cmdPath = append(cmdPath, args...)
	logger.Info("Command:", cmdPath)
	cmd = exec.Command(cmdPath[0], cmdPath[1:]...)
}
