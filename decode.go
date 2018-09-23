package qrcode

import (
	"fmt"
	"io"

	"github.com/tuotoo/qrcode"
)

type Decoder struct {
	r io.Reader
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (dec *Decoder) Decode(v interface{}) error {
	d, err := qrcode.Decode(dec.r)
	if err != nil {
		return err
	}
	switch v := v.(type) {
	case *string:
		*v = d.Content
	case io.Writer:
		_, err = v.Write([]byte(d.Content))
	default:
		return fmt.Errorf("not implemented")
	}
	return nil
}
