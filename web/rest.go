package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HandlerFunc rest handler
type HandlerFunc func(*gin.Context) (interface{}, error)

//Rest rest handle
func Rest(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if v, e := h(c); e == nil {
			c.JSON(http.StatusOK, v)
		} else {
			c.String(http.StatusInternalServerError, e.Error())
		}

	}
}
