package platform

import (
	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	"github.com/itpkg/chaos/i18n"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type Engine struct {
	I18n              *i18n.I18n      `inject:""`
	Dao               *Dao            `inject:""`
	Jwt               *Jwt            `inject:""`
	Logger            *logging.Logger `inject:""`
	Cache             *web.Cache      `inject:""`
	Oauth2GoogleState string          `inject:"oauth2.google.state"`
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
		&inject.Object{Value: "cache", Name: "cache.prefix"},
		&inject.Object{Value: "ga2", Name: "oauth2.google.state"},
	)

}
func (p *Engine) Migrate(db *gorm.DB) {
	i18n.Migrate(db)

	db.AutoMigrate(
		&Setting{}, &Notice{},
		&User{}, &Role{}, &Permission{}, &Log{},
	)

	db.Model(&User{}).AddUniqueIndex("idx_user_provider_type_id", "provider_type", "provider_id")
	db.Model(&Role{}).AddUniqueIndex("idx_roles_name_resource_type_id", "name", "resource_type", "resource_id")
	db.Model(&Permission{}).AddUniqueIndex("idx_permissions_user_role", "user_id", "role_id")
}
func (p *Engine) Seed() {}

func init() {
	web.Register(&Engine{})
}
