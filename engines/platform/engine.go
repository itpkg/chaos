package platform

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/ugorji/go/codec"
)

type Engine struct {
	Dao    *Dao            `inject:""`
	Logger *logging.Logger `inject:""`
}

func (p *Engine) Map(inj *inject.Graph) error {
	db, err := OpenDatabase()
	if err != nil {
		return err
	}

	enc, err := NewAesHmacEncryptor(Secret(120, 32), Secret(210, 32))
	if err != nil {
		return err
	}

	var hnd codec.MsgpackHandle

	return inj.Provide(
		&inject.Object{Value: db},
		&inject.Object{Value: OpenRedis()},
		&inject.Object{Value: enc},
		&inject.Object{Value: &Coder{Handle: &hnd}},
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
func (p *Engine) Seed()   {}
func (p *Engine) Worker() {}

func init() {
	web.Register(&Engine{})
}
