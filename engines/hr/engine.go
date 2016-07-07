package hr

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

//Engine engine
type Engine struct {
}

//Map map objects
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

//Mount mount web points
func (p *Engine) Mount(*gin.Engine) {

}

//Migrate db:migrate
func (p *Engine) Migrate(*gorm.DB) {}

//Seed db:seed
func (p *Engine) Seed() {}

//Worker worker
func (p *Engine) Worker() {}

//Shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
