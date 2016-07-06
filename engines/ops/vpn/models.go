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

	Begin time.Time `gorm:"not null;default:current_date;type:date" json:"begin"`
	End   time.Time `gorm:"not null;default:'1000-1-1';type:date" json:"end"`

	Online bool `gorm:"not null" json:"online"`
	Enable bool `gorm:"not null" json:"enable"`
}

//TableName table's name of User
func (User) TableName() string {
	return "vpn_users"
}

//Log vpn log
type Log struct {
	ID     uint `gorm:"primary_key" json:"id"`
	UserID uint `gorm:"not null" json:"-"`
	User   User `json:"-"`

	TrustedIp   string    `gorm:"type:VARCHAR(32)" json:"trusted_ip"`
	TrustedPort string    `gorm:"type:VARCHAR(16)" json:"trusted_port"`
	RemoteIp    string    `gorm:"type:VARCHAR(32)" json:"remote_ip"`
	RemotePort  string    `gorm:"type:VARCHAR(16)" json:"remote_port"`
	Start       time.Time `gorm:"not null;default:current_date;type:date" json:"start"`
	End         time.Time `gorm:"not null;default:'1000-1-1';type:date" json:"end"`
	Received    float64   `gorm:"not null;default:0" json:"received"`
	Send        float64   `gorm:"not null;default:0" json:"send"`
}

//TableName table's name of Log
func (Log) TableName() string {
	return "vpn_logs"
}
