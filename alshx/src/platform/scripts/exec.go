package scripts

import (
	"alshx/src/platform/logging"
	"os"
	"os/exec"
	"path/filepath"

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
	config.Entrypoint = filepath.FromSlash(config.Entrypoint)

	logger.Info("=== Script Details ===")
	logger.Info("Language:", config.Language)
	logger.Info("Entrypoint:", config.Entrypoint)
	logger.Info("Command:", config.Entrypoint)

	var cmd *exec.Cmd

	if config.Language == "go" {
		cmd = exec.Command("go", "run", config.Entrypoint)
	} else if config.Language == "ts-node" {
		installNodeModules(logger, folderPath)
		cmd = exec.Command("ts-node", append([]string{config.Entrypoint}, config.Args...)...)
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
