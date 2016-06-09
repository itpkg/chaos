package reading

import (
	"github.com/facebookgo/inject"
	"github.com/itpkg/chaos/engines/platform"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

type Engine struct {
	Db  *gorm.DB      `inject:""`
	Jwt *platform.Jwt `inject:""`
}

func (p *Engine) Map(*inject.Graph) error {
	return nil
}

func (p *Engine) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Note{})
}

func (p *Engine) Seed() {}

func (p *Engine) Worker() {}

func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
