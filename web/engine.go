package web

import (
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

//Engine web engine
type Engine interface {
	Map(*inject.Graph) error
	Mount(*mux.Router)
	Migrate(*gorm.DB)
	Seed()
	Worker()
	Shell() []cli.Command
}

var engines []Engine

//Register register engine
func Register(ens ...Engine) {
	engines = append(engines, ens...)
}

//Loop loop engines
func Loop(fn func(Engine) error) error {
	for _, en := range engines {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
