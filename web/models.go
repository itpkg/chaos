package web

import (
	"time"
)

//Model base model
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//Link link model
type Link struct {
	Href  string `json:"href"`
	Label string `json:"label"`
}
