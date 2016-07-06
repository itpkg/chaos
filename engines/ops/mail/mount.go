package mail

import (
	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
)

//Mount mount web
func (p *Engine) Mount(r *gin.Engine) {
	rg := r.Group("/ops/mail", p.Jwt.MustRolesHandler("ops"))
	rg.GET("/domains", web.Rest(p.indexDomain))
	rg.POST("/domains", web.Rest(p.createDomain))
	rg.DELETE("/domains/:id", web.Rest(p.deleteDomain))

	rg.GET("/users", web.Rest(p.indexUser))
	rg.POST("/users", web.Rest(p.createUser))
	rg.DELETE("/users/:id", web.Rest(p.deleteUser))

	rg.GET("/aliases", web.Rest(p.indexAlias))
	rg.POST("/aliases", web.Rest(p.createAlias))
	rg.DELETE("/aliases/:id", web.Rest(p.deleteAlias))
}
