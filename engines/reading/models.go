package reading

import (
	"github.com/itpkg/chaos/engines/platform"
	"github.com/itpkg/chaos/web"
)

//Note note model
type Note struct {
	web.Model

	UserID uint          `gorm:"not null" json:"user_id"`
	User   platform.User `json:"-"`

	Title string `gorm:"not null" json:"title"`
	Body  string `gorm:"not null;type:text" json:"body"`
	Share bool   `gorm:"not null" json:"share"`
}

//TableName table's name of Note
func (Note) TableName() string {
	return "reading_notes"
}
