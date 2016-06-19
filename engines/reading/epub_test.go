package reading_test

import (
	"io/ioutil"
	"testing"

	"github.com/chonglou/epubgo"
)

func TestEpub(t *testing.T) {
	b, e := epubgo.Open("tmp/N28n0015.epub")
	if e != nil {
		t.Fatal(e)
	}
	defer b.Close()

	t.Logf("%+v", b)

	t.Log("========")
	for _, v := range []string{"title", "subject", "publisher", "creator", "date"} {
		if s, e := b.Metadata(v); e == nil {
			t.Logf("%s = %v", v, s)
		} else {
			t.Fatal(e)
		}
	}

	t.Logf("============")
	for it, err := b.Navigation(); !it.IsLast(); it.Next() {
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s: %s", it.Title(), it.URL())
	}

	t.Logf("============")
	for it, err := b.Spine(); !it.IsLast(); it.Next() {
		if err != nil {
			t.Fatal(err)
		}
		pg, err := it.Open()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("========= %s ========", it.URL())

		if buf, err := ioutil.ReadAll(pg); err == nil {
			t.Log(string(buf))
		} else {
			t.Fatal(err)
		}
		pg.Close()
	}

}
