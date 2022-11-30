package handlers

import (
	"fmt"

	"github.com/alshdavid/alshx/kit/arguments"
	"github.com/alshdavid/alshx/tools/envm/platform/meta"
)

func HandlerVersion() arguments.RouteHandler {
	return func(args *arguments.Arguments) {
		fmt.Println("", meta.Version)
	}
}
