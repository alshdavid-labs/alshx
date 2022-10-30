package handler_version

import (
	"github.com/alshdavid/alshx/tools/random/platform/colors"
	"github.com/alshdavid/alshx/tools/random/platform/info"
)

func Handler() {
	colors.PrintLog(
		"Random Generation Tool",
		"Version: "+info.Version,
	)
}
