package qrcode

import (
	"fmt"
	"io"

	qrcode "github.com/skip2/go-qrcode"
)

type Encoder struct {
	w io.Writer
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (enc *Encoder) Encode(v interface{}) error {
	var str string
	switch v := v.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return fmt.Errorf("not implemented")
	}
	q, err := qrcode.New(str, qrcode.Highest)
	if err != nil {
		return err
	}
	const size = 512
	d, err := q.PNG(size)
	if err != nil {
		return err
	}
	_, err = enc.w.Write(d)
	return err
}
