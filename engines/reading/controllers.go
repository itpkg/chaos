package reading

import (
	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/engines/platform"
	"github.com/itpkg/chaos/web"
)

func (p *Engine) index(c *gin.Context) (interface{}, error) {
	var notes []Note
	db := p.Db
	if o, ok := c.Get("user"); ok {
		db = db.Where("user_id = ? or share", o.(*platform.User).ID)
	} else {
		db = db.Where("share")
	}
	err := db.Order("updated_at DESC").Find(&notes).Error
	return notes, err
}

//NoteFm note form model
type NoteFm struct {
	ID    uint   `form:"id"`
	Title string `form:"title" binding:"required"`
	Body  string `form:"body" binding:"required"`
}

func (p *Engine) create(c *gin.Context) (interface{}, error) {
	u := c.MustGet("user").(*platform.User)
	var fm NoteFm
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	note := Note{UserID: u.ID, Body: fm.Body, Title: fm.Title}
	err := p.Db.Create(&note).Error
	return note, err
}
func (p *Engine) update(c *gin.Context) (interface{}, error) {
	u := c.MustGet("user").(*platform.User)
	id := c.Param("id")
	var note Note
	if err := p.Db.Where("id = ? AND user_id = ?", id, u.ID).First(&note).Error; err != nil {
		return nil, err
	}
	var fm NoteFm
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	err := p.Db.Model(&note).Updates(map[string]interface{}{
		"title": fm.Title,
		"body":  fm.Body,
	}).Error
	return note, err

}
func (p *Engine) show(c *gin.Context) (interface{}, error) {
	u := c.MustGet("user").(*platform.User)
	id := c.Param("id")
	var note Note
	err := p.Db.Where("id = ? AND user_id = ?", id, u.ID).First(&note).Error
	return note, err
}

func (p *Engine) delete(c *gin.Context) (interface{}, error) {

	u := c.MustGet("user").(*platform.User)
	id := c.Param("id")
	err := p.Db.Where("id = ? AND user_id = ?", id, u.ID).Delete(Note{}).Error
	return web.OK, err
}

//Mount mount router
func (p *Engine) Mount(r *gin.Engine) {
	r.GET("/reading/notes", p.Jwt.CurrentUserHandler(false), web.Rest(p.index))
	g := r.Group("/reading", p.Jwt.CurrentUserHandler(true))
	g.POST("/notes", web.Rest(p.create))
	g.GET("/notes/:id", web.Rest(p.show))
	g.POST("/notes/:id", web.Rest(p.update))
	g.DELETE("/notes/:id", web.Rest(p.delete))
}
