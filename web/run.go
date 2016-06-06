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
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{}
	for _, en := range engines {
		cmd := en.Shell()
		app.Commands = append(app.Commands, cmd...)
	}
	return app.Run(os.Args)
}
