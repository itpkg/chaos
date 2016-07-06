package mail_test

import (
	"testing"

	"github.com/itpkg/chaos/engines/ops/mail"
)

func TestSha(t *testing.T) {
	en := mail.Encryptor{}
	pln := "hello"
	enc, err := en.Sum(pln, 32)
	if err != nil {
		t.Fatal(err)
	}
	if !en.Chk(pln, enc) {
		t.Fatal("error on check")
	}
	t.Logf("doveadm pw -t {SSHA512}%s -p %s", enc, pln)
}
