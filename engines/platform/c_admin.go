package platform

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
)

func (p *Engine) getAdminSiteInfo(c *gin.Context) {
	lng := c.MustGet("locale").(string)
	ifo := p._siteInfoMap(lng)
	c.JSON(http.StatusOK, ifo)
}

func (p *Engine) postAdminSiteInfo(c *gin.Context) (interface{}, error) {
	c.Request.ParseForm()

	lng := c.MustGet("locale").(string)
	for _, k := range []string{"title", "subTitle", "keywords", "description",
		"aboutUs", "copyright"} {
		if err := p.Dao.Set(p._siteKey(lng, k), c.Request.Form.Get(k), false); err != nil {
			return nil, err
		}
	}
	for _, k := range []string{"name", "email"} {
		if err := p.Dao.Set(p._siteAuthorKey(k), c.Request.Form.Get("author"+strings.Title(k)), false); err != nil {
			return nil, err
		}
	}
	var links []web.Link
	if err := json.Unmarshal([]byte(c.Request.Form.Get("navLinks")), &links); err != nil {
		return nil, err
	}
	if err := p.Dao.Set(p._siteKey("", "navLinks"), links, false); err != nil {
		return nil, err
	}
	return web.OK, nil
}

func (p *Engine) deleteAdminCache(c *gin.Context) (interface{}, error) {
	err := p.Cache.Flush()
	return web.OK, err
}
