package platform

import (
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/SermoDigital/jose/jws"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/satori/go.uuid"
)

type Dao struct {
	Db        *gorm.DB        `inject:""`
	Encryptor Encryptor       `inject:""`
	Logger    *logging.Logger `inject:""`
}

func (p *Dao) Set(k string, v interface{}, f bool) error {
	buf, err := msgpack.Marshal(v)
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
		//err = p.Db.Save(&m).Error
		err = p.Db.Model(&m).Updates(map[string]interface{}{
			"flag": f,
			"val":  buf,
		}).Error
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
	return msgpack.Unmarshal(m.Val, v)
}

//-----------------------------------------------------------------------------

func (p *Dao) UserClaims(u *User, days int) jws.Claims {
	cm := jws.Claims{}
	now := time.Now()
	cm.SetNotBefore(now)
	cm.SetExpiration(now.AddDate(0, 0, days))
	cm.SetSubject(u.Name)
	cm.Set("uid", u.UID)

	cm.Set("roles", p.Authority(u.ID, "-", 0))
	return cm
}

func (p *Dao) AddUser(pid, pty, email, name, home, logo string) (*User, error) {
	var u User
	var err error
	if p.Db.Where("provider_id = ? AND provider_type = ?", pid, pty).First(&u).RecordNotFound() {
		u.Email = email
		u.Name = name
		u.Logo = logo
		u.Home = home
		u.UID = uuid.NewV4().String()
		u.ProviderID = pid
		u.ProviderType = pty
		now := time.Now()
		u.ConfirmedAt = &now
		u.SignInCount = 1
		u.LastSignIn = &now
		err = p.Db.Create(&u).Error
	} else {
		err = p.Db.Model(&u).Updates(map[string]interface{}{
			"email":         email,
			"name":          name,
			"logo":          logo,
			"home":          home,
			"sign_in_count": u.SignInCount + 1,
		}).Error
	}
	return &u, err
}

func (p *Dao) GetUser(uid string) (*User, error) {
	var u User
	err := p.Db.Where("uid = ?", uid).First(&u).Error
	return &u, err
}

func (p *Dao) Authority(user uint, rty string, rid uint) []string {
	var items []Role
	if err := p.Db.
		Where("resource_type = ? AND resource_id = ?", rty, rid).
		Find(&items).Error; err != nil {
		p.Logger.Error(err)
	}
	var roles []string
	for _, r := range items {
		var pm Permission
		if err := p.Db.
			Where("role_id = ? AND user_id = ? ", r.ID, user).
			First(&pm).Error; err != nil {
			p.Logger.Error(err)
			continue
		}
		if pm.Enable() {
			roles = append(roles, r.Name)
		}
	}

	return roles
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

func (p *Dao) Allow(role uint, user uint, years, months, days int) error {
	begin := time.Now()
	end := begin.AddDate(years, months, days)
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
