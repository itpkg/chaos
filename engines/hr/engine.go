package hr

import (
	"github.com/itpkg/chaos/web"
	"github.com/urfave/cli"
)

type Engine struct {
}

func (p *Engine) Mount() {

}
func (p *Engine) Migrate() {}
func (p *Engine) Seed()    {}
func (p *Engine) Worker()  {}
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
