package ops

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

type Engine struct {
}

func (p *Engine) Map(*inject.Graph) error {
	return nil
}
func (p *Engine) Mount(*gin.Engine) {

}

func (p *Engine) Migrate(*gorm.DB) {}
func (p *Engine) Seed()            {}
func (p *Engine) Worker()          {}
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
