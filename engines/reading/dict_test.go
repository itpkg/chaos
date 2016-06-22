package reading_test

import (
	"os/exec"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	items := map[string]bool{
		"aaa":     true,
		"中文测试":    true,
		"a中文a":    true,
		"_bbb":    true,
		"bbb_":    true,
		"b123bb_": true,

		"": false,
		"123456789012345678901234567890123": false,
		"wer'we":   false,
		"wer$fwe":  false,
		"awwe sfe": false,
	}
	s := `^[\p{Han}\w]{1,32}$`
	exp := regexp.MustCompile(s) //("^[a-z\u4e00-\u9fa5]$") //(`^[\w\u4e00-\u9fa5]+$`)

	for k, v := range items {
		// r, e := regexp.MatchString(s, k)
		// t.Logf("check %s %v %v %v", k, v, r, e)
		if exp.MatchString(k) == v {
			t.Logf("check %s %v: passed", k, v)
		} else {
			t.Fatalf("check %s %v: NOT passed", k, v)
		}
	}

}
func TestDict(t *testing.T) {
	name := "hello"
	out, err := exec.Command("sdcv", "--data-dir", "tmp/dict", name).Output()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("query[%s], %s", name, out)
}
