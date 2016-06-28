package cms

import (
	"github.com/facebookgo/inject"
	"github.com/itpkg/chaos/web"
	"github.com/urfave/cli"
)

//Engine engine
type Engine struct {
}

//Map map
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

//Seed db:seed
func (p *Engine) Seed() {}

//Worker worker
func (p *Engine) Worker() {}

//Shell shell
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
