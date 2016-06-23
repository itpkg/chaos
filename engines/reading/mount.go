package reading

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
)

//Mount mount router
func (p *Engine) Mount(r *gin.Engine) {

	gf := r.Group("/reading", p.Jwt.CurrentUserHandler(false))
	gf.GET("/blogs", p.Cache.Page(time.Hour*24, web.Rest(p.indexBlogs)))
	gf.GET("/blog/*name", p.Cache.Page(time.Hour*24, p.showBlog))
	gf.GET("/books", p.Cache.Page(time.Hour*24, web.Rest(p.indexBooks)))
	gf.GET("/book/:id/*name", p.Cache.Page(time.Hour*24, p.showBook))
	gf.GET("/dict", p.getDict)
	gf.POST("/dict", p.postDict)

	gt := r.Group("/reading", p.Jwt.CurrentUserHandler(true))
	gt.DELETE("/books/:id", p.Jwt.MustAdminHandler(), web.Rest(p.deleteBook))
	gt.GET("/notes", web.Rest(p.indexNotes))
	gt.POST("/notes", web.Rest(p.createNote))
	gt.GET("/notes/:id", web.Rest(p.showNote))
	gt.POST("/notes/:id", web.Rest(p.updateNote))
	gt.DELETE("/notes/:id", web.Rest(p.deleteNote))
}
