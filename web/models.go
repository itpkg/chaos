package web

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Link struct {
	Href  string `json:"href"`
	Label string `json:"label"`
}

var OK = gin.H{"ok": true}
