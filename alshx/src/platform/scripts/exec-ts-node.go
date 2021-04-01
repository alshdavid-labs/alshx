package scripts

import (
	"alshx/src/platform/logging"
	"os"
	"os/exec"
)

func execTsNode(
	logger *logging.Logger,
	cmd *exec.Cmd,
	config *Meta,
	args []string,
	folderPath string,
) *exec.Cmd {
	if !hasCommand("node") || !hasCommand("npx") {
		logger.Log("Node is not installed")
		os.Exit(1)
	}
	if !hasCommand("yarn") {
		logger.Log("Yarn is not installed")
		os.Exit(1)
	}
	installNodeModules(logger, folderPath)
	cmdPath := []string{"npx", "ts-node", config.Entrypoint}
	cmdPath = append(cmdPath, config.Args...)
	cmdPath = append(cmdPath, args...)
	logger.Info("Command:", cmdPath)
	cmd = exec.Command(cmdPath[0], cmdPath[1:]...)
	return cmd
}
