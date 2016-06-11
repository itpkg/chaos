package platform

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
)

func (p *Engine) Mount(r *gin.Engine) {
	ag := r.Group("/admin", p.Jwt.CurrentUserHandler(true), p.Jwt.MustAdminHandler())
	ag.GET("/site/info", p.getAdminSiteInfo)
	ag.POST("/site/info", web.Rest(p.postAdminSiteInfo))
	ag.DELETE("/cache", web.Rest(p.deleteAdminCache))
	ag.GET("/notices", web.Rest(p.getNotices))
	ag.POST("/notices", web.Rest(p.postNotices))
	ag.DELETE("/notices/:id", web.Rest(p.deleteNotice))

	r.GET("/notices", p.Cache.Page(time.Hour*24, web.Rest(p.getNotices)))

	r.GET("/personal/self", p.Jwt.CurrentUserHandler(true), web.Rest(p.getPersonalSelf))
	r.GET("/personal/logs", p.Jwt.CurrentUserHandler(true), web.Rest(p.getPersonalLogs))
	r.DELETE("/personal/signOut", p.Jwt.CurrentUserHandler(true), p.deleteSignOut)

	r.GET("/locales/:lang", p.Cache.Page(time.Hour*24, p.getLocale))
	r.GET("/site/info", p.Cache.Page(time.Hour*24, p.getSiteInfo))

	r.POST("/oauth2/callback", web.Rest(p.postOauth2Callback))
}
