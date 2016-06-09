package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*gin.Context) (interface{}, error)

func Rest(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if v, e := h(c); e == nil {
			c.JSON(http.StatusOK, v)
		} else {
			c.String(http.StatusInternalServerError, e.Error())
		}

	}
}
