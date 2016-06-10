package i18n

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func LocaleHandler(c *gin.Context) {
	// written := false
	// 1. Check URL arguments.
	lng := c.Request.URL.Query().Get("locale")

	// 2. Get language information from cookies.
	if len(lng) == 0 {
		if ck, er := c.Request.Cookie("locale"); er == nil {
			lng = ck.Value
		}
	}
	// else {
	// 	written = true
	// }

	// 3. Get language information from 'Accept-Language'.
	if len(lng) == 0 {
		al := c.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			lng = al[:5]
		}
	}

	tag := Match(lng)

	// if written {
	// 	c.SetCookie("locale", tag.String(), 1<<31-1, "/", "", false, false)
	// }
	c.Set("locale", &tag)
}

func Match(lng string) language.Tag {
	tag, _, _ := matcher.Match(language.Make(lng))
	return tag
}

var matcher language.Matcher

func init() {
	matcher = language.NewMatcher([]language.Tag{
		language.AmericanEnglish,
		language.SimplifiedChinese,
	})
}
