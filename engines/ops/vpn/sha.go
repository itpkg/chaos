package vpn

import (
	"crypto/sha512"
	"encoding/base64"

	"github.com/op/go-logging"
)

type Encryptor struct {
	Logger *logging.Logger `inject:""`
}

func (p *Encryptor) Sum(plain string) string {
	code := sha512.Sum512([]byte(plain))
	return base64.StdEncoding.EncodeToString(code[:])
}
