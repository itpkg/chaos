package hr

import (
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

type Engine struct {
}

func (p *Engine) Mount() {

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
