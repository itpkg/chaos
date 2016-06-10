package platform

import (
	"net/http"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type Jwt struct {
	Key    []byte               `inject:"jwt.key"`
	Method crypto.SigningMethod `inject:"jwt.method"`
	Logger *logging.Logger      `inject:""`
	Db     *gorm.DB             `inject:""`
}

func (p *Jwt) Validate(buf []byte) (jwt.Claims, error) {
	tk, err := jws.ParseJWT(buf)
	if err != nil {
		return nil, err
	}
	err = tk.Validate(p.Key, p.Method)
	return tk.Claims(), err
}

func (p *Jwt) Handler(c *gin.Context) {
	tkn, err := jws.ParseFromRequest(c.Request, jws.Compact)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	if err := tkn.Verify(p.Key, p.Method); err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}
	var user User
	data := tkn.Payload().(map[string]interface{})
	if err := p.Db.Where("uid = ?", data["uid"]).First(&user).Error; err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}
	if !user.IsAvailable() {
		c.String(http.StatusForbidden, "bad user status")
		c.Abort()
		return
	}
	c.Set("user", &user)
}

func (p *Jwt) Sum(cm jws.Claims) ([]byte, error) {
	jt := jws.NewJWT(cm, p.Method)
	return jt.Serialize(p.Key)
}
