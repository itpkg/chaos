package vpn

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/engines/platform"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

//Engine engine
type Engine struct {
	Db        *gorm.DB        `inject:""`
	Logger    *logging.Logger `inject:""`
	Jwt       *platform.Jwt   `inject:""`
	Encryptor *Encryptor      `inject:""`
}

//Map map object
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

//Migrate db:migrate
func (p *Engine) Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Log{})
}

//Seed db:seed
func (p *Engine) Seed() {}

//Worker worker
func (p *Engine) Worker() {}

//Mount mount web
func (p *Engine) Mount(r *gin.Engine) {
	rg := r.Group("/ops/vpn", p.Jwt.MustRolesHandler("ops"))
	rg.GET("/users", web.Rest(p.indexUser))
	rg.POST("/users", web.Rest(p.createUser))
	rg.DELETE("/users/:id", web.Rest(p.deleteUser))
}

func init() {
	web.Register(&Engine{})
}
