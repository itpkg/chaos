package platform

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/garyburd/redigo/redis"
	"github.com/op/go-logging"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
)

//Jwt jwt helper
type Jwt struct {
	Render *render.Render       `inject:""`
	Key    []byte               `inject:"jwt.key"`
	Method crypto.SigningMethod `inject:"jwt.method"`
	Logger *logging.Logger      `inject:""`
	Redis  *redis.Pool          `inject:""`
	Dao    *Dao                 `inject:""`
}

//Validate check jwt
func (p *Jwt) Validate(buf []byte) (jwt.Claims, error) {
	tk, err := jws.ParseJWT(buf)
	if err != nil {
		return nil, err
	}
	err = tk.Validate(p.Key, p.Method)
	return tk.Claims(), err
}

//MustAdmin check must have admin role
func (p *Jwt) MustAdmin(req *http.Request, fn http.HandlerFunc) http.HandlerFunc {
	if p.IsAdmin(req) {
		fn
	}
	return errors.New("must have admin role")
}

//MustAdminHandler check must have admin role
func (p *Jwt) IsAdmin(req *http.Request) bool {
	return p.HasRoles(req, "admin")
}

//MustRolesHandler check must have one roles at least
func (p *Jwt) HasRoles(req *http.Request, roles ...string) bool {
	u, e := p.CurrentUser(req)
	if e == nil {
		for _, a := range p.Dao.Authority(u.ID, "-", 0) {
			for _, r := range roles {
				if a == r {
					return true
				}
			}
		}
	}
	return false

}

//CurrentUserHandler inject current user
func (p *Jwt) CurrentUser(req *http.Request) (*User, error) {
	tkn, err := jws.ParseFromRequest(req, jws.Compact)
	if err != nil {
		return nil, err
	}

	if err = tkn.Verify(p.Key, p.Method); err != nil {
		return nil, err
	}
	var user User
	data := tkn.Payload().(map[string]interface{})
	if err = p.Dao.Db.Where("uid = ?", data["uid"]).First(&user).Error; err != nil {
		return nil, err
	}
	if !user.IsAvailable() {
		return nil, errors.New("bad user status")
	}
	return &user, nil
}

func (p *Jwt) key(kid string) string {
	return fmt.Sprintf("token://%s", kid)
}

//Sum create jwt token
func (p *Jwt) Sum(cm jws.Claims, days int) ([]byte, error) {
	kid := uuid.NewV4()
	now := time.Now()
	cm.SetNotBefore(now)
	cm.SetExpiration(now.AddDate(0, 0, days))
	cm.Set("kid", kid)
	//TODO using kid

	jt := jws.NewJWT(cm, p.Method)
	return jt.Serialize(p.Key)
}
