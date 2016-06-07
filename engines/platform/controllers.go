package platform

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Engine) Mount(r *gin.Engine) {
	r.GET("/info", func(c *gin.Context) {
		ifo := make(map[string]interface{})
		for _, v := range []string{"title", "subTitle", "description", "copyright"} {
			ifo[v] = "site." + v
		}
		ifo["author"] = map[string]string{
			"name":  "username",
			"email": "aaa@aaa.com",
		}

		c.JSON(http.StatusOK, ifo)
	})
}
