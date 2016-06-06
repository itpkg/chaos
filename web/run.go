package web

import (
	"os"

	"github.com/urfave/cli"
)

func Run() error {
	app := cli.NewApp()
	app.Name = "chaos"
	app.Version = "v20160606"
	app.Usage = "it-package web application."
	return app.Run(os.Args)
}
