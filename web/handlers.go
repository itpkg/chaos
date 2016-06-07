package web

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func LocaleHandler(c *gin.Context) {
	// 1. Check URL arguments.
	lng := c.Request.URL.Query().Get("locale")

	// 2. Get language information from cookies.
	if len(lng) == 0 {
		if ck, er := c.Request.Cookie("locale"); er == nil {
			lng = ck.String()
		}
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lng) == 0 {
		al := c.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			lng = al[:5]
		}
	}

	tag, _, _ := matcher.Match(language.Make(lng))
	c.Set("locale", &tag)
}

var matcher language.Matcher

func init() {
	matcher = language.NewMatcher([]language.Tag{
		language.AmericanEnglish,
		language.SimplifiedChinese,
	})
}
