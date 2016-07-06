package mail

import "github.com/itpkg/chaos/web"

//Domain mail transport
type Domain struct {
	web.Model
	Name    string `gorm:"type:varchar(128);unique;not null"`
	Users   []User
	Aliases []Alias
}

//TableName table's name of Domain
func (Domain) TableName() string {
	return "mail_domains"
}

//User mail user
type User struct {
	web.Model

	DomainID uint `gorm:"not null"`
	Domain   Domain

	Email    string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(128);not null"`
	Name     string `gorm:"type:varchar(128);not null;index"`
	Home     string `gorm:"type:varchar(128);not null"`
}

//TableName table's name of User
func (User) TableName() string {
	return "mail_users"
}

//Alias alias
type Alias struct {
	web.Model

	DomainID uint `gorm:"not null"`
	Domain   Domain

	Source      string `gorm:"type:varchar(255);unique;primary_key;not null"`
	Destination string `gorm:"type:varchar(255);not null"`
}

//TableName table's name of Alias
func (Alias) TableName() string {
	return "mail_aliases"
}
