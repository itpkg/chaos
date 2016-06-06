package platform

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type Engine struct {
	Dao        *Dao            `inject:""`
	Logger     *logging.Logger `inject:""`
	MailSender *MailSender     `inject:""`
}

func (p *Engine) Map(inj *inject.Graph) error {
	db, err := web.OpenDatabase()
	if err != nil {
		return err
	}

	enc, err := NewAesHmacEncryptor(Secret(120, 32), Secret(210, 32))
	if err != nil {
		return err
	}

	return inj.Provide(
		&inject.Object{Value: db},
		&inject.Object{Value: web.OpenRedis()},
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
func (p *Engine) Worker(srv *machinery.Server) {
	srv.RegisterTask("email", p.MailSender.Send)
}

func init() {
	web.Register(&Engine{})
}
