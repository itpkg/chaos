package mail

import "github.com/itpkg/chaos/web"

//Domain mail transport
type Domain struct {
	web.Model

	Name    string  `gorm:"type:varchar(128);unique;not null" json:"name"`
	Users   []User  `json:"users"`
	Aliases []Alias `json:"aliases"`
}

//TableName table's name of Domain
func (Domain) TableName() string {
	return "mail_domains"
}

//User mail user
type User struct {
	web.Model

	DomainID uint   `gorm:"not null" json:"-"`
	Domain   Domain `json:"domain"`

	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Name     string `gorm:"type:varchar(128);not null;index" json:"name"`
}

//TableName table's name of User
func (User) TableName() string {
	return "mail_users"
}

//Alias alias
type Alias struct {
	web.Model

	DomainID uint   `gorm:"not null" json:"-"`
	Domain   Domain `json:"domain"`

	Source      string `gorm:"type:varchar(255);unique;not null" json:"source"`
	Destination string `gorm:"type:varchar(255);not null;index" json:"destination"`
}

//TableName table's name of Alias
func (Alias) TableName() string {
	return "mail_aliases"
}
