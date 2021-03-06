package reading

import (
	"github.com/facebookgo/inject"
	"github.com/itpkg/chaos/engines/platform"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

//Engine Reading engine
type Engine struct {
	Db     *gorm.DB        `inject:""`
	Jwt    *platform.Jwt   `inject:""`
	Logger *logging.Logger `inject:""`
	Cache  *web.Cache      `inject:""`
}

//Map mapping objects
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

//Migrate db:migrate
func (p *Engine) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Note{}, &Book{})
}

//Seed db:seed
func (p *Engine) Seed() {}

//Worker worker
func (p *Engine) Worker() {}

func init() {
	web.Register(&Engine{})
}
