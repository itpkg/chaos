package mail

import "github.com/gin-gonic/gin"

//DomainFm form for transport
type DomainFm struct {
	Name string `form:"name" binding:"required"`
}

func (p *Engine) indexDomain(c *gin.Context) (interface{}, error) {
	var items []Domain
	err := p.Db.Order("name DESC").Find(&items).Error
	return items, err
}

func (p *Engine) createDomain(c *gin.Context) (interface{}, error) {
	var fm DomainFm
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	d := &Domain{Name: fm.Name}
	err := p.Db.Create(d).Error
	return d, err
}

func (p *Engine) deleteDomain(c *gin.Context) (interface{}, error) {
	var d Domain
	err := p.Db.Where("id = ?", c.Param("id")).First(&d).Error
	if err == nil {
		err = p.Db.Delete(&d).Error
	}
	return d, err
}
