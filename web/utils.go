package web

import (
	"math/rand"
	"time"
)

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
