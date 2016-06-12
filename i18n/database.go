package i18n

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"golang.org/x/text/language"
)

//Migrate migrate database
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Locale{})
	db.Model(&Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")
}

//Locale locale model
type Locale struct {
	ID        uint   `gorm:"primary_key"`
	Lang      string `gorm:"not null;type:varchar(8);index"`
	Code      string `gorm:"not null;index;type:VARCHAR(255)"`
	Message   string `gorm:"not null;type:varchar(800)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//DatabaseProvider db provider
type DatabaseProvider struct {
	Db     *gorm.DB        `inject:""`
	Logger *logging.Logger `inject:""`
}

//Set set locale
func (p *DatabaseProvider) Set(lng *language.Tag, code, message string) {
	var l Locale
	var err error
	if p.Db.Where("lang = ? AND code = ?", lng.String(), code).First(&l).RecordNotFound() {
		l.Lang = lng.String()
		l.Code = code
		l.Message = message
		err = p.Db.Create(&l).Error
	} else {
		l.Message = message
		err = p.Db.Save(&l).Error
	}
	if err != nil {
		p.Logger.Error(err)
	}
}

//Get get locale
func (p *DatabaseProvider) Get(lng *language.Tag, code string) string {
	var l Locale
	if err := p.Db.Where("lang = ? AND code = ?", lng.String(), code).First(&l).Error; err != nil {
		p.Logger.Error(err)
	}
	return l.Message

}

//Del del locale
func (p *DatabaseProvider) Del(lng *language.Tag, code string) {
	if err := p.Db.Where("lang = ? AND code = ?", lng.String(), code).Delete(Locale{}).Error; err != nil {
		p.Logger.Error(err)
	}
}

//Keys list locale keys
func (p *DatabaseProvider) Keys(lng *language.Tag) ([]string, error) {
	var keys []string
	err := p.Db.Model(&Locale{}).Where("lang = ?", lng.String()).Pluck("code", &keys).Error

	return keys, err
}
