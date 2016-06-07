package platform

import (
	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type Engine struct {
	Dao    *Dao            `inject:""`
	Jwt    *Jwt            `inject:""`
	Logger *logging.Logger `inject:""`
}

func (p *Engine) Map(inj *inject.Graph) error {

	enc, err := NewAesHmacEncryptor(Secret(120, 32), Secret(210, 32))
	if err != nil {
		return err
	}

	return inj.Provide(
		&inject.Object{Value: enc},
		&inject.Object{Value: Secret(320, 32), Name: "jwt.key"},
		&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
	)

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

func init() {
	web.Register(&Engine{})
}
