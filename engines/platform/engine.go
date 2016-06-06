package platform

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/jrallison/go-workers"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

type Engine struct {
	Dao    *Dao            `inject:""`
	Logger *logging.Logger `inject:""`
}

func (p *Engine) Map(inj *inject.Graph) error {

	enc, err := NewAesHmacEncryptor(Secret(120, 32), Secret(210, 32))
	if err != nil {
		return err
	}

	return inj.Provide(
		&inject.Object{Value: enc},
	)

}
func (p *Engine) Mount(*gin.Engine) {

}

func (p *Engine) Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&Setting{}, &Locale{}, &Notice{},
		&User{}, &Role{}, &Permission{}, &Log{},
	)
	db.Model(&Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")
	db.Model(&User{}).AddUniqueIndex("idx_user_provider_type_id", "provider_type", "provider_id")
	db.Model(&Role{}).AddUniqueIndex("idx_roles_name_resource_type_id", "name", "resource_type", "resource_id")
	db.Model(&Permission{}).AddUniqueIndex("idx_permissions_user_role", "user_id", "role_id")
}
func (p *Engine) Seed() {}
func (p *Engine) Worker() {
	workers.Process("email", func(msg *workers.Msg) {

		p.Logger.Infof("GET JOB %s@email", msg.Jid())
		p.Logger.Debugf("ARGS: %+v", msg.Args())

	}, viper.GetInt("workers.email"))
}

func init() {
	web.Register(&Engine{})
}
