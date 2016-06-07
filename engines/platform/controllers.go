package platform

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func (p *Engine) Mount(r *gin.Engine) {
	r.GET("/info", func(c *gin.Context) {
		lng := c.MustGet("locale").(*language.Tag)
		ifo := make(map[string]interface{})
		for _, v := range []string{"title", "subTitle", "description", "copyright"} {
			ifo[v] = "site." + v
		}
		ifo["author"] = map[string]string{
			"name":  "username",
			"email": "aaa@aaa.com",
		}
		ifo["lang"] = lng.String()

		c.JSON(http.StatusOK, ifo)
	})
}
