package auth

import "github.com/itpkg/chaos/web"

type Engine struct {
}

func (p *Engine) Mount() {

}
func (p *Engine) Migrate() {}
func (p *Engine) Seed()    {}
func (p *Engine) Worker()  {}
func (p *Engine) Shell()   {}

func init() {
	web.Register(&Engine{})
}
