package platform

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
)

func (p *Engine) deleteSignOut(c *gin.Context) {
	u := c.MustGet("user").(*User)
	p.Dao.Log(u.ID, "sign out")
	c.JSON(http.StatusOK, web.OK)
}

func (p *Engine) getPersonalLogs(c *gin.Context) (interface{}, error) {
	u := c.MustGet("user").(*User)
	var logs []Log
	err := p.Dao.Db.
		Select([]string{"created_at", "message"}).
		Where("user_id = ?", u.ID).
		Order("ID DESC").Limit(200).
		Find(&logs).Error
	return logs, err
}

func (p *Engine) getPersonalSelf(c *gin.Context) (interface{}, error) {
	return c.MustGet("user").(*User), nil
}
