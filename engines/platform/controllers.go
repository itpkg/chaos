package platform

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/i18n"
	"github.com/itpkg/chaos/web"
	"golang.org/x/oauth2"
)

func (p *Engine) getSiteInfo(c *gin.Context) {
	lng := c.MustGet("locale").(string)
	ifo := p._siteInfoMap(lng)

	var gcf oauth2.Config
	if err := p.Dao.Get("google.oauth", &gcf); err != nil {
		p.Logger.Error(err)
	}
	ifo["oauth2"] = map[string]string{
		"google": gcf.AuthCodeURL(p.Oauth2GoogleState),
	}

	c.JSON(http.StatusOK, ifo)
}

func (p *Engine) getLocale(c *gin.Context) {
	lng := i18n.Match(c.Param("lang"))
	c.JSON(http.StatusOK, p.I18n.Items(lng.String()))
}

func (p *Engine) _siteKey(lng, key string) string {
	return fmt.Sprintf("%s://site/%s", lng, key)
}
func (p *Engine) _siteAuthorKey(key string) string {
	return p._siteKey("", "author/"+key)
}

func (p *Engine) _siteInfoMap(lng string) map[string]interface{} {
	ifo := make(map[string]interface{})
	for _, k := range []string{
		"title", "subTitle", "keywords",
		"description", "copyright", "aboutUs"} {
		var v string
		if e := p.Dao.Get(p._siteKey(lng, k), &v); e == nil {
			ifo[k] = v
		} else {
			p.Logger.Error(e)
			ifo[k] = fmt.Sprintf("%s.site.%s", lng, k)
		}
	}

	author := make(map[string]string)
	for _, k := range []string{"name", "email"} {
		var v string
		if e := p.Dao.Get(p._siteAuthorKey(k), &v); e == nil {
			author[k] = v
		} else {
			p.Logger.Error(e)
			author[k] = fmt.Sprintf("site.author.%s", k)
		}
	}
	ifo["author"] = author

	var links []web.Link
	if err := p.Dao.Get(p._siteKey("", "navLinks"), &links); err != nil {
		p.Logger.Error(err)
		links = append(
			links,
			web.Link{Label: "platform.pages.home", Href: "/home"},
			web.Link{Label: "platform.pages.about_us", Href: "/about-us"},
		)
	}
	ifo["navLinks"] = links

	return ifo
}
