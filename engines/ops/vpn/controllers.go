package vpn

import (
	"errors"

	"github.com/gin-gonic/gin"
)

//UserFm form for user
type UserFm struct {
	Name       string `form:"name" binding:"required"`
	Email      string `form:"email" binding:"required"`
	Password   string `form:"password" binding:"required"`
	RePassword string `form:"re_password" binding:"required"`
}

func (p *Engine) indexUser(c *gin.Context) (interface{}, error) {
	var items []User
	err := p.Db.Order("name DESC").Find(&items).Error
	return items, err
}

func (p *Engine) createUser(c *gin.Context) (interface{}, error) {
	var fm UserFm
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	if fm.Password != fm.RePassword {
		return nil, errors.New("passwords not match")
	}

	user := &User{
		Email:    fm.Email,
		Password: p.Encryptor.Sum(fm.Password),
		Name:     fm.Name,
	}
	err := p.Db.Create(user).Error
	return user, err
}

func (p *Engine) deleteUser(c *gin.Context) (interface{}, error) {
	var u User
	err := p.Db.Where("id = ?", c.Param("id")).First(&u).Error
	if err == nil {
		err = p.Db.Delete(&u).Error
	}
	return u, err
}
