package web

import (
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

type Engine interface {
	Mount()
	Migrate(*gorm.DB)
	Seed()
	Worker()
	Shell() []cli.Command
}

var engines []Engine

func Register(ens ...Engine) {
	engines = append(engines, ens...)
}

func Loop(fn func(Engine) error) error {
	for _, en := range engines {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
