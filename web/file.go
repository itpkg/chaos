package web

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

//Bytes file render
func Bytes(name string, buf []byte) gin.HandlerFunc {
	return func(c *gin.Context) {

		ext := path.Ext(name)
		switch {
		case ext == ".xhtml" || ext == ".html" || ext == ".htm":
			c.Data(http.StatusOK, "text/html; charset=UTF-8", buf)
		case ext == ".css":
			c.Data(http.StatusOK, "text/css; charset=UTF-8", buf)
		default:
			c.String(http.StatusNotFound, fmt.Sprintf("bad format %s", ext))
		}
	}
}
