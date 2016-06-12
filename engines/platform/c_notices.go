package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
)

//NoticeFm form for notice
type NoticeFm struct {
	Content string `form:"content" binding:"required"`
}

func (p *Engine) getNotices(c *gin.Context) (interface{}, error) {
	var ns []Notice
	err := p.Dao.Db.
		Select([]string{"id", "content", "created_at"}).
		Order("ID DESC").Limit(120).Find(&ns).Error
	return ns, err
}

func (p *Engine) postNotices(c *gin.Context) (interface{}, error) {
	lng := c.MustGet("locale").(string)
	var fm NoticeFm
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	n := Notice{Content: fm.Content, Lang: lng}
	err := p.Dao.Db.Create(&n).Error
	return n, err
}

func (p *Engine) deleteNotice(c *gin.Context) (interface{}, error) {
	id := c.Param("id")
	err := p.Dao.Db.Where("id = ?", id).Delete(&Notice{}).Error
	return web.OK, err
}
