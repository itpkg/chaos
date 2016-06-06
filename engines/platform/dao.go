package platform

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"golang.org/x/text/language"
)

type Dao struct {
	Db        *gorm.DB        `inject:""`
	Encryptor Encryptor       `inject:""`
	Coder     *Coder          `inject:""`
	Logger    *logging.Logger `inject:""`
}

func (p *Dao) Set(k string, v interface{}, f bool) error {
	buf, err := p.Coder.To(v)
	if err != nil {
		return err
	}
	if f {
		buf, err = p.Encryptor.Encode(buf)
		if err != nil {
			return err
		}
	}
	var m Setting
	null := p.Db.Where("key = ?", k).First(&m).RecordNotFound()
	m.Key = k
	m.Val = buf
	m.Flag = f
	if null {
		err = p.Db.Create(&m).Error
	} else {
		err = p.Db.Save(&m).Error
	}
	return err
}

func (p *Dao) Get(k string, v interface{}) error {
	var m Setting
	err := p.Db.Where("key = ?", k).First(&m).Error
	if err != nil {
		return err
	}
	if m.Flag {
		if m.Val, err = p.Encryptor.Decode(m.Val); err != nil {
			return err
		}
	}
	return p.Coder.From(m.Val, v)
}

//-----------------------------------------------------------------------------

func (p *Dao) SetLocale(lng *language.Tag, code, message string) {
	var l Locale
	var err error
	if p.Db.
		Where("lang = ? AND code = ?", lng.String(), code).
		First(&l).RecordNotFound() {
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

func (p *Dao) T(lng *language.Tag, code string) string {
	var l Locale
	if err := p.Db.
		Where("lang = ? AND code = ?", lng.String(), code).
		First(&l).Error; err != nil {
		p.Logger.Error(err)
	}
	return l.Message

}

func (p *Dao) DelLocale(lng *language.Tag, code string) {
	if err := p.Db.
		Where("lang = ? AND code = ?", lng.String(), code).
		Delete(Locale{}).Error; err != nil {
		p.Logger.Error(err)
	}
}

func (p *Dao) GetLocaleKeys(lng *language.Tag) []string {
	var keys []string
	if err := p.Db.
		Model(&Locale{}).
		Where("lang = ?", lng.String()).
		Pluck("code", &keys).Error; err != nil {
		p.Logger.Error(err)
	}
	return keys
}

//-----------------------------------------------------------------------------

func (p *Dao) GetUser(uid string) (*User, error) {
	var u User
	err := p.Db.Where("uid = ?", uid).First(&u).Error
	return &u, err
}

func (p *Dao) Is(user uint, name string) bool {
	return p.Can(user, name, "-", 0)
}

func (p *Dao) Can(user uint, name string, rty string, rid uint) bool {
	var r Role
	if p.Db.
		Where("name = ? AND resource_type = ? AND resource_id = ?", name, rty, rid).
		First(&r).
		RecordNotFound() {
		return false
	}
	var pm Permission
	if p.Db.
		Where("user_id = ? AND role_id = ?", user, r.ID).
		First(&pm).
		RecordNotFound() {
		return false
	}

	return pm.Enable()
}

func (p *Dao) Role(name string, rty string, rid uint) (*Role, error) {
	var e error
	r := Role{}
	db := p.Db
	if db.
		Where("name = ? AND resource_type = ? AND resource_id = ?", name, rty, rid).
		First(&r).
		RecordNotFound() {
		r = Role{
			Name:         name,
			ResourceType: rty,
			ResourceID:   rid,
		}
		e = db.Create(&r).Error

	}
	return &r, e
}

func (p *Dao) Deny(role uint, user uint) error {
	return p.Db.
		Where("role_id = ? AND user_id = ?", role, user).
		Delete(Permission{}).Error
}

func (p *Dao) Allow(role uint, user uint, dur time.Duration) error {
	begin := time.Now()
	end := begin.Add(dur)
	var count int
	p.Db.
		Model(&Permission{}).
		Where("role_id = ? AND user_id = ?", role, user).
		Count(&count)
	if count == 0 {
		return p.Db.Create(&Permission{
			UserID: user,
			RoleID: role,
			Begin:  begin,
			End:    end,
		}).Error
	}
	return p.Db.
		Model(&Permission{}).
		Where("role_id = ? AND user_id = ?", role, user).
		UpdateColumns(map[string]interface{}{"begin": begin, "end": end}).Error

}
