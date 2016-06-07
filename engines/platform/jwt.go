package platform

import (
	"fmt"
	"net/http"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/gin-gonic/gin"
)

type Jwt struct {
	Key    []byte               `inject:"jwt.key"`
	Method crypto.SigningMethod `inject:"jwt.method"`
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
	js, err := jws.ParseFromRequest(c.Request, jws.Flat)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	if !js.IsJWT() {
		c.AbortWithError(
			http.StatusInternalServerError,
			fmt.Errorf("http request header not have jwt"),
		)
	}
	//TODO
}

func (p *Jwt) Sum(cm jws.Claims) ([]byte, error) {
	jt := jws.NewJWT(cm, p.Method)
	return jt.Serialize(p.Key)
}
