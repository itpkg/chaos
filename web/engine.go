package web

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

type Engine interface {
	Map(*inject.Graph) error
	Mount(*gin.Engine)
	Migrate(*gorm.DB)
	Seed()
	Worker(*machinery.Server)
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
