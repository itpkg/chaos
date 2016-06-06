package web

import "github.com/urfave/cli"

type Engine interface {
	Mount()
	Migrate()
	Seed()
	Worker()
	Shell() []cli.Command
}

var engines []Engine

func Register(ens ...Engine) {
	engines = append(engines, ens...)
}
