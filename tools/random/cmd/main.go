package main

import (
	"os"

	handler_guid "github.com/alshdavid/alshx/tools/random/cmd/handlers/guid"
	handler_help "github.com/alshdavid/alshx/tools/random/cmd/handlers/help"
	handler_number "github.com/alshdavid/alshx/tools/random/cmd/handlers/number"
	handler_string "github.com/alshdavid/alshx/tools/random/cmd/handlers/string"
	handler_version "github.com/alshdavid/alshx/tools/random/cmd/handlers/version"
	"github.com/alshdavid/alshx/tools/random/platform/args"
)

var Routes = struct {
	Empty   string
	Number  string
	String  string
	Guid    string
	Version string
	Help    string
}{
	Empty:   "",
	Number:  "number",
	String:  "string",
	Guid:    "guid",
	Version: "version",
	Help:    "help",
}

func main() {
	arguments := args.NewArgs(os.Args[1:])

	if arguments.Command == Routes.Empty {
		arguments.Command = Routes.Help
	}

	switch arguments.Command {
	case Routes.Version:
		handler_version.Handler()
	case Routes.Help:
		handler_help.Handler()
	case Routes.Number:
		handler_number.Handler(arguments.Args)
	case Routes.String:
		handler_string.Handler(arguments.Args)
	case Routes.Guid:
		handler_guid.Handler()
	}
}
