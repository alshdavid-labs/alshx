package main

import (
	"os"

	"github.com/alshdavid/alshx/kit/arguments"
	"github.com/alshdavid/alshx/tools/envm/cmd/handlers"
)

func main() {
	app := arguments.NewRouter()

	app.
		Route("version", "v").
		Handle(handlers.HandlerVersion())

	app.Run(os.Args)
}
