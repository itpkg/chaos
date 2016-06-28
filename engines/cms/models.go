package cms

import (
	"github.com/itpkg/chaos/engines/platform"
	"github.com/itpkg/chaos/web"
	"github.com/jinzhu/gorm"
)

//Article article
type Article struct {
	web.Model
	Title   string        `gorm:"not null;index;size:255" json:"title"`
	Summary string        `gorm:"not null;size:500" json:"summary"`
	Body    string        `gorm:"not null;type:text" json:"body"`
	UserID  uint          `gorm:"not null" json:"user_id"`
	User    platform.User `json:"-"`
	Tags    []Tag         `json:"tags"`
}

//TableName table's name
func (Article) TableName() string {
	return "cms_articles"
}

//Tag tag
type Tag struct {
	web.Model
	Name     string    `gorm:"not null;index;size:255" json:"name"`
	Lang     string    `gorm:"not null;size:8;index" json:"name"`
	Articles []Article `json:"articles"`
}

//TableName table's name
func (Tag) TableName() string {
	return "cms_tags"
}

//Comment comment
type Comment struct {
	web.Model
	Body      string        `gorm:"not null;type:text" json:"body"`
	UserID    uint          `gorm:"not null" json:"user_id"`
	User      platform.User `json:"-"`
	ArticleID uint          `gorm:"not null"`
	Article   Article       `json:"article"`
}

//TableName table's name
func (Comment) TableName() string {
	return "cms_comments"
}

//Attachment attachment
type Attachment struct {
	web.Model
	UID    string        `gorm:"not null;size:36" json:"uid"`
	Name   string        `gorm:"not null;size:255" json:"name"`
	Ext    string        `gorm:"not null;size:5" json:"ext"`
	Size   uint          `gorm:"not null" json:"size"`
	UserID uint          `gorm:"not null" json:"user_id"`
	User   platform.User `json:"-"`
}

//TableName table's name
func (Attachment) TableName() string {
	return "cms_attachments"
}

//-----------------------------------------------------------------------------

//Migrate db:migrate
func (p *Engine) Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&Article{}, &Tag{}, &Comment{}, &Attachment{},
	)

	db.Model(&Tag{}).AddUniqueIndex("idx_cms_tag_name_lang", "name", "lang")

}
