package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/alshdavid/alshx/tools/archive_desktop/platform/flags"
	"github.com/alshdavid/alshx/tools/archive_desktop/platform/logging"
	"github.com/alshdavid/alshx/tools/archive_desktop/platform/paths"
)

func main() {
	var name = paths.ArchiveNamePath()
	var desktopPath = paths.GetDesktopPath()
	var archivePath = paths.GetArchivePath()
	var destPath = filepath.Join(archivePath, name)
	var args = flags.NewArgs()
	var logger = logging.NewLogger(args.Verbose)

	if args.Version {
		logger.Log("Archive desktop, probably the latest verison")
		os.Exit(0)
	}

	if args.Dry {
		logger.Info("INFO: Dry Mode - No files will be moved")
	}
	logger.Info("INFO: Source", desktopPath)

	files, _ := ioutil.ReadDir(desktopPath)

	if runtime.GOOS == "windows" {
		publicPath := paths.GetPublicDesktop()
		logger.Info("INFO: Source", publicPath)
		moreFiles, _ := ioutil.ReadDir(publicPath)
		files = append(files, moreFiles...)
	}

	logger.Info("INFO: Destination", destPath)

	logger.Info("")

	for _, file := range files {
		fileName := file.Name()
		filePath := filepath.Join(desktopPath, fileName)

		if runtime.GOOS == "windows" {
			if fileName == "desktop.ini" {
				logger.Info("INFO: Skipping", fileName)
				continue
			}
			if strings.HasSuffix(fileName, ".lnk") {
				logger.Info("INFO: Deleting", fileName)
				if !args.Dry {
					os.RemoveAll(filePath)
				}
				continue
			}
		} else if runtime.GOOS == "darwin" {
			if fileName == ".DS_Store" || fileName == ".localized" {
				logger.Info("INFO: Skipping", fileName)
				continue
			}
		}
		logger.Info("INFO: Copying", fileName)
		if !args.Dry {
			os.MkdirAll(destPath, 0755)
			os.Rename(filePath, filepath.Join(destPath, fileName))
		}
	}
}
