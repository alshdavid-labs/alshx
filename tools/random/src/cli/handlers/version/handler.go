package handler_version

import (
	"random/src/platform/colors"
	"random/src/platform/info"
)

func Handler() {
	colors.PrintLog(
		"Random Generation Tool",
		"Version: "+info.Version,
	)
}
