package platform

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Engine) Mount(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		SendMail("aaa@aaa.com", "subject", "<h1>body</h1>", true, "/tmp/aaa.txt", "/tmp/bbb.txt")
		c.JSON(http.StatusOK, "hello")
	})
}
