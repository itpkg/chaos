package platform_test

import (
	"testing"
	"time"

	"github.com/itpkg/chaos/engines/platform"
	"github.com/ugorji/go/codec"
)

type S struct {
	S string
	I int
	T time.Time
}

func TestCoder(t *testing.T) {
	var hnd codec.MsgpackHandle
	cod := platform.Coder{Handle: &hnd}
	s1 := S{S: "hello, test!", I: 2016, T: time.Now()}
	buf, err := cod.To(&s1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MSGPACK(%+v)=%s", s1, buf)
	var s2 S
	if err = cod.From(buf, &s2); err != nil {
		t.Fatal(err)
	}
	if s1.I != s2.I || s1.S != s2.S {
		t.Fatalf("want %+v, get %+v", s1, s2)
	}
}
