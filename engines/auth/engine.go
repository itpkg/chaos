package auth

import (
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
)

type Engine struct {
}

func (p *Engine) Mount() {

}
func (p *Engine) Migrate(*gorm.DB) {}
func (p *Engine) Seed()            {}
func (p *Engine) Worker()          {}

func init() {
	web.Register(&Engine{})
}
