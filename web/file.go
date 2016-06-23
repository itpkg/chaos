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
		case ext == ".md" || ext == ".txt":
			c.Data(http.StatusOK, "text/plain; charset=UTF-8", buf)
		case ext == ".html" || ext == ".htm":
			c.Data(http.StatusOK, "text/html; charset=UTF-8", buf)
		case ext == ".xhtml":
			c.Data(http.StatusOK, "application/xhtml+xml; charset=UTF-8", buf)
		case ext == ".xml":
			c.Data(http.StatusOK, "text/xml; charset=UTF-8", buf)
		case ext == ".css":
			c.Data(http.StatusOK, "text/css; charset=UTF-8", buf)
		case ext == ".png" || ext == ".jpg" || ext == ".jpeg":
			c.Data(http.StatusOK, fmt.Sprintf("image/%s", ext), buf)
		default:
			c.String(http.StatusNotFound, fmt.Sprintf("bad format %s", ext))
		}
	}
}
