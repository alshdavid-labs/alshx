package main

import (
	"os"
	handler_guid "random/src/cli/handlers/guid"
	handler_help "random/src/cli/handlers/help"
	handler_number "random/src/cli/handlers/number"
	handler_string "random/src/cli/handlers/string"
	handler_version "random/src/cli/handlers/version"
	"random/src/platform/args"
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
