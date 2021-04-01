package scripts

import (
	"alshx/src/platform/files"
	"alshx/src/platform/logging"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func Exec(
	logger *logging.Logger,
	folderPath string,
	cacheFolderPath string,
	name string,
	args []string,
) {
	metaFilePath := filepath.Join(folderPath, "meta.yaml")
	newMetaFilePath := filepath.Join(cacheFolderPath, "meta.yaml")

	logger.Info("MetaPath:", metaFilePath)
	logger.Info("MetaCachePath:", newMetaFilePath)

	meta := loadConfig(metaFilePath)
	newerMeta := loadConfig(newMetaFilePath)

	if newerMeta.Version > meta.Version {
		logger.Info("Updating script with newer version", meta.Version, ">", newerMeta.Version)
		os.RemoveAll(folderPath)
		files.CopyDir(cacheFolderPath, folderPath)
	}

	logger.Info("Language:", meta.Language)
	if meta.Entrypoint != "" {
		logger.Info("Entrypoint:", meta.Entrypoint)
	}
	if meta.Command != "" {
		logger.Info("Command:", meta.Command)
	}

	var cmd *exec.Cmd

	if meta.Language == "go" {
		cmd = execGo(logger, cmd, meta, args)
	} else if meta.Language == "ts-node" {
		cmd = execTsNode(logger, cmd, meta, args, folderPath)
	} else if meta.Language == "node" {
		cmd = execNode(logger, cmd, meta, args, folderPath)
	} else if meta.Language == "shell" {
		cmd = execShell(logger, cmd, meta, args, folderPath)
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

func loadConfig(path string) *Meta {
	config := &Meta{Args: []string{}}
	configBytes, _ := os.ReadFile(path)
	yaml.UnmarshalStrict(configBytes, config)
	config.Entrypoint = filepath.FromSlash(config.Entrypoint)
	return config
}
