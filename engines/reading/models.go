package reading

import (
	"github.com/itpkg/chaos/engines/platform"
	"github.com/itpkg/chaos/web"
)

type Note struct {
	web.Model

	UserID uint          `gorm:"not null" json:"-"`
	User   platform.User `json:"-"`

	Title string `gorm:"not null" json:"title"`
	Body  string `gorm:"not null;type:text" json:"body"`
}

func (Note) TableName() string {
	return "reading_notes"
}
