package scripts

import (
	"alshx/src/platform/files"
	"alshx/src/platform/logging"
	"os"
	"os/exec"
	"path/filepath"
)

func execNode(
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
	cmdPath := []string{"node", config.Entrypoint}
	cmdPath = append(cmdPath, config.Args...)
	cmdPath = append(cmdPath, args...)
	logger.Info("Command:", cmdPath)
	cmd = exec.Command(cmdPath[0], cmdPath[1:]...)
	return cmd
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
