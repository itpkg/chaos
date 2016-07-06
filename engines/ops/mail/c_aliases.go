package mail

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//AliasFm form for alias
type AliasFm struct {
	Source      string `form:"source" binding:"required"`
	Domain      uint   `form:"domain" binding:"required"`
	Destination string `form:"destination" binding:"required"`
}

func (p *Engine) indexAlias(c *gin.Context) (interface{}, error) {
	var items []Alias
	err := p.Db.Order("source DESC").Find(&items).Error
	return items, err
}

func (p *Engine) createAlias(c *gin.Context) (interface{}, error) {
	var fm AliasFm
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var d Domain
	if err := p.Db.Where("id = ?", fm.Domain).First(&d).Error; err != nil {
		return nil, err
	}
	a := &Alias{
		Source:      fmt.Sprintf("%s@%s", fm.Source, d.Name),
		Destination: fmt.Sprintf("%s@%s", fm.Destination, d.Name),
	}
	err := p.Db.Create(a).Error
	return a, err
}

func (p *Engine) deleteAlias(c *gin.Context) (interface{}, error) {
	var a Alias
	err := p.Db.Where("id = ?", c.Param("id")).First(&a).Error
	if err == nil {
		err = p.Db.Delete(&a).Error
	}
	return a, err
}
