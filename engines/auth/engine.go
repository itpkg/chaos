package auth

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
)

type Engine struct {
}

func (p *Engine) Map(inj *inject.Graph) error {
	db, err := OpenDatabase()
	if err != nil {
		return err
	}
	return inj.Provide(&inject.Object{Value: db}, &inject.Object{Value: OpenRedis()})

}
func (p *Engine) Mount(*gin.Engine) {

}

func (p *Engine) Migrate(*gorm.DB) {}
func (p *Engine) Seed()            {}
func (p *Engine) Worker()          {}

func init() {
	web.Register(&Engine{})
}
