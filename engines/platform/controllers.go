package platform

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func (p *Engine) info(c *gin.Context) {

	lng := c.MustGet("locale").(*language.Tag).String()
	ifo := make(map[string]interface{})
	ifo["lang"] = lng
	for _, k := range []string{"title", "subTitle", "description", "copyright"} {
		var v string
		if e := p.Dao.Get(fmt.Sprintf("%s://site/%s", lng, k), &v); e == nil {
			ifo[k] = v
		} else {
			p.Logger.Error(e)
		}
	}
	for _, k := range []string{"name", "email"} {
		var v string
		if e := p.Dao.Get(fmt.Sprintf("%s://site/author/%s", lng, k), &v); e == nil {
			ifo[k] = v
		} else {
			p.Logger.Error(e)
		}
	}

	c.JSON(http.StatusOK, ifo)
}

func (p *Engine) Mount(r *gin.Engine) {
	r.GET("/info", p.Cache.Page(time.Hour*24, p.info))
}
