package mail

import (
	"crypto/sha512"
	"encoding/base64"
	"math/rand"

	"github.com/op/go-logging"
)

type Encryptor struct {
	Logger *logging.Logger `inject:""`
}

func (p *Encryptor) Sum(plain string, size int) (string, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return p.sum(plain, salt), nil

}
func (p *Encryptor) sum(plain string, salt []byte) string {
	buf := append([]byte(plain), salt...)
	code := sha512.Sum512(buf)
	return base64.StdEncoding.EncodeToString(append(code[:], salt...))
}
func (p *Encryptor) Chk(plain string, code string) bool {
	buf, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		p.Logger.Error(err)
		return false
	}
	if len(buf) <= sha512.Size {
		return false
	}
	return code == p.sum(plain, buf[sha512.Size:])
}
