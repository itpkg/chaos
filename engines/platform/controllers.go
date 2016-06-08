package platform

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
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
			ifo[k] = fmt.Sprintf("%s.site.%s", lng, k)
		}
	}

	author := make(map[string]string)
	for _, k := range []string{"name", "email"} {
		var v string
		if e := p.Dao.Get(fmt.Sprintf("site/author/%s", k), &v); e == nil {
			author[k] = v
		} else {
			p.Logger.Error(e)
			author[k] = fmt.Sprintf("site.author.%s", k)
		}
	}
	ifo["author"] = author

	var links []Link
	if err := p.Dao.Get("site/links", &links); err != nil {
		p.Logger.Error(err)
	}
	ifo["links"] = links

	var gcf oauth2.Config
	if err := p.Dao.Get("google.oauth", &gcf); err != nil {
		p.Logger.Error(err)
	}
	ifo["oauth"] = map[string]string{
		"google": gcf.AuthCodeURL("ga2"),
	}

	c.JSON(http.StatusOK, ifo)
}

func (p *Engine) google(c *gin.Context) {
	c.String(http.StatusInternalServerError, "fuck")
}

func (p *Engine) Mount(r *gin.Engine) {
	r.GET("/info", p.Cache.Page(time.Hour*24, p.info))
	r.POST("/oauth2/callback", p.google)
}
