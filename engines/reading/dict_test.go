package reading_test

import (
	"os/exec"
	"testing"
)

func TestDict(t *testing.T) {
	name := "hello"
	out, err := exec.Command("sdcv", "--data-dir", "tmp/dict", name).Output()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("query[%s], %s", name, out)
}
