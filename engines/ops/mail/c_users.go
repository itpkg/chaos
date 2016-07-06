package mail

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

//UserFm form for user
type UserFm struct {
	Uid        string `form:"uid" binding:"required"`
	Domain     uint   `form:"domain" binding:"required"`
	Name       string `form:"name" binding:"required"`
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
	pwd, err := p.Encryptor.Sum(fm.Password, 32)
	if err != nil {
		return nil, err
	}

	var d Domain
	if err = p.Db.Where("id = ?", fm.Domain).Limit(1).Find(&d).Error; err != nil {
		return nil, err
	}

	user := &User{
		Email:    fmt.Sprintf("%s@%s", fm.Uid, d.Name),
		Password: pwd,
		Name:     fm.Name,
	}
	err = p.Db.Create(user).Error
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
