package reading_test

import (
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

	t.Log("==== Metadata ====")
	flds := b.MetadataFields()
	//t.Logf("metadatas: %v", flds)
	for _, v := range flds {
		if s, e := b.Metadata(v); e == nil {
			t.Logf("%s = %v", v, s)
		} else {
			t.Fatal(e)
		}
	}

	t.Logf("==== Navigation ========")
	for it, err := b.Navigation(); !it.IsLast(); it.Next() {
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s: %s", it.Title(), it.URL())
	}

	// t.Logf("==== Open file ====")
	// if fd, err := b.OpenFile("stylesheet.css"); err == nil {
	// 	defer fd.Close()
	// 	if buf, err := ioutil.ReadAll(fd); err == nil {
	// 		t.Log(string(buf))
	// 	} else {
	// 		t.Fatal(err)
	// 	}
	// } else {
	// 	t.Fatal(err)
	// }

	t.Logf("===== Spine =======")
	for it, err := b.Spine(); !it.IsLast(); it.Next() {
		if err != nil {
			t.Fatal(err)
		}
		pg, err := it.Open()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("========= %s ========", it.URL())

		// if buf, err := ioutil.ReadAll(pg); err == nil {
		// 	t.Log(string(buf))
		// } else {
		// 	t.Fatal(err)
		// }
		pg.Close()
	}

}
