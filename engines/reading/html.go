package reading

import (
	"encoding/xml"
	"io"
)

//HTML model
type HTML struct {
	Body Body `xml:"body"`
}

//Body model
type Body struct {
	Content string `xml:",innerxml"`
}

//Parse parse body from html file
func Parse(r io.Reader) (string, error) {
	var ht HTML
	err := xml.NewDecoder(r).Decode(&ht)
	return ht.Body.Content, err
}
