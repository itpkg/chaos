package vpn

import (
	"time"

	"github.com/itpkg/chaos/web"
)

//User vpn user
type User struct {
	web.Model

	Name     string `gorm:"type:varchar(128);index;not null" json:"name"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`

	Begin time.Time `gorm:"not null;default:current_date;type:date"`
	End   time.Time `gorm:"not null;default:'1000-1-1';type:date"`
}

//TableName table's name of User
func (User) TableName() string {
	return "vpn_users"
}
