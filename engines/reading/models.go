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

	Title string `gorm:"not null;index" json:"title"`
	Body  string `gorm:"not null;type:text" json:"body"`
	Share bool   `gorm:"not null" json:"share"`
}

//TableName table's name of Note
func (Note) TableName() string {
	return "reading_notes"
}

//Book book model
type Book struct {
	web.Model

	Name      string `sql:"not null;unique_index" json:"name"`
	Title     string `sql:"not null;index" json:"title"`
	Creator   string `sql:"not null;index" json:"creator"`
	Subject   string `sql:"not null;index" json:"subject"`
	Publisher string `sql:"not null;index" json:"publisher"`
	Version   string `sql:"not null;index" json:"version"`
	Home      string `sql:"not null;index" json:"home"`
}

//TableName table name
func (Book) TableName() string {
	return "reading_books"
}
