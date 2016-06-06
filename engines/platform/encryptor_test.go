package platform_test

import (
	"encoding/hex"
	"testing"

	"github.com/itpkg/chaos/engines/platform"
)

func TestEncryptor(t *testing.T) {
	en, er := platform.NewAesHmacEncryptor([]byte("12345678901234567890123456789012"), []byte("hello"))
	if er != nil {
		t.Fatal(er)
	}
	s1 := []byte("hello, test!")
	testAes(t, en, s1)
	testHmac(t, en, s1)
}

func testHmac(t *testing.T, en platform.Encryptor, s1 []byte) {
	b1 := en.Sum(s1)
	b2 := en.Sum(s1)
	t.Logf("HMAC(%s)=%s", s1, hex.EncodeToString(b1))
	t.Logf("HMAC(%s)=%s", s1, hex.EncodeToString(b2))
	if !en.Equal(s1, b1) {
		t.Fatalf("error on check")
	}
}

func testAes(t *testing.T, en platform.Encryptor, s1 []byte) {
	if buf, err := en.Encode(s1); err == nil {
		t.Logf("AES(%s)=%s", s1, hex.EncodeToString(buf))
		if s2, err := en.Decode(buf); err == nil {
			t.Logf("get %s", s2)
			if string(s2) != string(s1) {
				t.Fatalf("want %s, get %s", s1, s2)
			}
		} else {
			t.Fatal(err)
		}
	} else {
		t.Fatal(err)
	}
}
