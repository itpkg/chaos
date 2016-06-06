package platform

import (
	"math/rand"
	"time"

	"github.com/jrallison/go-workers"
	"github.com/spf13/viper"
)

func Secret(i, l int) []byte {
	secret := viper.GetString("secrets")
	return []byte(secret[i : i+l])
}

func SendMail(to, subject, body string, html bool, files ...string) {
	workers.Enqueue("email", "send", []interface{}{to, subject, body, html, files})
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
