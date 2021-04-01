package scripts

import (
	"alshx/src/platform/logging"
	"os/exec"
)

func execTsNode(
	logger *logging.Logger,
	cmd *exec.Cmd,
	config *Meta,
	args []string,
	folderPath string,
) {
	if !hasCommand("node") || !hasCommand("npx") {
		logger.Log("Node is not installed")
		return
	}
	if !hasCommand("yarn") {
		logger.Log("Yarn is not installed")
		return
	}
	installNodeModules(logger, folderPath)
	cmdPath := []string{"npx", "ts-node", config.Entrypoint}
	cmdPath = append(cmdPath, config.Args...)
	cmdPath = append(cmdPath, args...)
	logger.Info("Command:", cmdPath)
	cmd = exec.Command(cmdPath[0], cmdPath[1:]...)
}
