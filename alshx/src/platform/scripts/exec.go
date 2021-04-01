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
	metaPath := filepath.Join(folderPath, "meta.yaml")
	logger.Info("MetaPath:", metaPath)

	config := &Meta{Args: []string{}}
	configBytes, _ := os.ReadFile(metaPath)
	yaml.UnmarshalStrict(configBytes, config)
	config.Entrypoint = filepath.FromSlash(config.Entrypoint)

	logger.Info("Language:", config.Language)
	if config.Entrypoint != "" {
		logger.Info("Entrypoint:", config.Entrypoint)
	}
	if config.Command != "" {
		logger.Info("Command:", config.Command)
	}

	var cmd *exec.Cmd

	if config.Language == "go" {
		cmd = exec.Command("go", "run", config.Entrypoint)
	} else if config.Language == "ts-node" {
		installNodeModules(logger, folderPath)
		cmdPath := append([]string{"npx", "ts-node", config.Entrypoint}, config.Args...)
		logger.Info("Command:", cmdPath)
		cmd = exec.Command(cmdPath[0], cmdPath[1:]...)
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
