package platform

import (
	"math/rand"
	"time"

	"github.com/spf13/viper"
)

func Secret(i, l int) []byte {
	secret := viper.GetString("secrets")
	return []byte(secret[i : i+l])
}

func RandomStr(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	buf := make([]rune, n)
	for i := range buf {
		buf[i] = letters[rand.Intn(len(letters))]
	}
	return string(buf)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
