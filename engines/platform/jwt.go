package platform

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/satori/go.uuid"
)

type Jwt struct {
	Key    []byte               `inject:"jwt.key"`
	Method crypto.SigningMethod `inject:"jwt.method"`
	Logger *logging.Logger      `inject:""`
	Db     *gorm.DB             `inject:""`
	Redis  *redis.Pool          `inject:""`
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

func (p *Jwt) key(kid string) string {
	return fmt.Sprintf("token://%s", kid)
}

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
