package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/i18n"
	"github.com/itpkg/chaos/web"
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

	var links []web.Link
	if err := p.Dao.Get("site/links", &links); err != nil {
		p.Logger.Error(err)
		links = append(
			links,
			web.Link{Label: "platform.pages.home", Href: "/home"},
			web.Link{Label: "platform.pages.about_us", Href: "/about-us"},
		)
	}
	ifo["links"] = links

	var gcf oauth2.Config
	if err := p.Dao.Get("google.oauth", &gcf); err != nil {
		p.Logger.Error(err)
	}
	ifo["oauth"] = map[string]string{
		"google": gcf.AuthCodeURL(p.Oauth2GoogleState),
	}

	c.JSON(http.StatusOK, ifo)
}

type OauthFm struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

func (p *Engine) oauthCallback(c *gin.Context) (interface{}, error) {
	var fm OauthFm
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	//p.Logger.Debugf("%+v", fm)
	var u *User
	var e error
	switch fm.State {
	case p.Oauth2GoogleState:
		u, e = p.google(fm.Code)
	default:
		e = errors.New("bad state")
	}

	var tk []byte
	if e == nil {
		tk, e = p.Jwt.Sum(p.Dao.UserClaims(u, 7))
	}
	return gin.H{"token": string(tk)}, e

}

func (p *Engine) google(code string) (*User, error) {
	var cfg oauth2.Config
	if err := p.Dao.Get("google.oauth", &cfg); err != nil {
		return nil, err
	}

	tok, err := cfg.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	cli := cfg.Client(oauth2.NoContext, tok)
	res, err := cli.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var gu GoogleUser
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&gu); err != nil {
		return nil, err
	}
	return p.Dao.AddUser(gu.ID, "google", gu.Email, gu.Name, gu.Link, gu.Picture)
}

func (p *Engine) locale(c *gin.Context) {
	lng := i18n.Match(c.Param("lang"))
	c.JSON(http.StatusOK, p.I18n.Items(&lng))
}

func (p *Engine) Mount(r *gin.Engine) {
	r.GET("/locales/:lang", p.Cache.Page(time.Hour*24, p.locale))
	r.GET("/info", p.Cache.Page(time.Hour*24, p.info))
	r.POST("/oauth2/callback", web.Rest(p.oauthCallback))
}
