package scripts

import (
	"alshx/src/platform/files"
	"alshx/src/platform/logging"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func Exec(
	logger *logging.Logger,
	folderPath string,
	name string,
	args []string,
) {
	config := &Meta{}
	configBytes, _ := os.ReadFile(filepath.Join(folderPath, "meta.yaml"))
	yaml.UnmarshalStrict(configBytes, config)
	config.Entrypoint = filepath.Join(strings.Split(config.Entrypoint, "/")...)

	logger.Info("=== Script Details ===")
	logger.Info("Language", config.Language)
	logger.Info("Action", config.Action)
	logger.Info("Entrypoint:", config.Entrypoint)
	logger.Info("Command:", config.Entrypoint)

	var cmd *exec.Cmd

	if config.Language == "go" {
		cmd = exec.Command("go", "run", config.Entrypoint)
	} else if config.Language == "commands" {
		if files.NotExists(filepath.Join(folderPath, "node_modules")) {
			logger.Log("Installing Node Modules")
			yarn := exec.Command("yarn")
			yarn.Dir = folderPath
			yarn.Run()
		}
		commandFull := strings.Split(config.Command, " ")
		if len(commandFull) == 1 {
			cmd = exec.Command(commandFull[0])
		} else {
			cmd = exec.Command(commandFull[0], commandFull[1:]...)
		}

	} else {
		logger.Log("Unable to process script")
		return
	}

	cmd.Dir = folderPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
