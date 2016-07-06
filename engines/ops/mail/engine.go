package mail

import (
	"github.com/facebookgo/inject"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/urfave/cli"
)

//Engine engine
type Engine struct {
	Db        *gorm.DB        `inject:""`
	Encryptor Encryptor       `inject:""`
	Logger    *logging.Logger `inject:""`
}

//Map map object
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

//Migrate db:migrate
func (p *Engine) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Domain{}, &User{}, &Alias{})
}

//Seed db:seed
func (p *Engine) Seed() {}

//Worker worker
func (p *Engine) Worker() {}

//Shell command
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
