package platform

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/jrallison/go-workers"
	"github.com/spf13/viper"
)

//Secret get key from secret config
func Secret(i, l int) []byte {
	secret := viper.GetString("secrets")
	return []byte(secret[i : i+l])
}

//SendMail add mail job to queue
func SendMail(to, subject, body string, html bool, files ...string) {
	workers.Enqueue("email", "send", []interface{}{to, subject, body, html, files})
}

//Shell run shell command
func Shell(cmd string, args ...string) error {
	bin, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}
	return syscall.Exec(bin, append([]string{cmd}, args...), os.Environ())
}
