package platform

import (
	"bytes"

	"github.com/ugorji/go/codec"
)

type Coder struct {
	Handle codec.Handle
}

func (p *Coder) To(o interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := codec.NewEncoder(&buf, p.Handle)
	err := enc.Encode(o)
	return buf.Bytes(), err
}

func (p *Coder) From(buf []byte, obj interface{}) error {
	dec := codec.NewDecoder(bytes.NewReader(buf), p.Handle)
	return dec.Decode(obj)
}
