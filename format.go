package qrcode

import (
	"fmt"
)

type Content struct {
	URL      string
	Email    string
	SSID     string
	Password string
	Type     string
	IsHidden bool
}

func (c Content) String() string {
	switch {
	case c.URL != "":
		return c.URLString()
	case c.Email != "":
		return c.EmailString()
	case c.SSID != "":
		return c.WiFiString()
	}
	return ""
}

func (c Content) URLString() string {
	return c.URL
}

func (c Content) EmailString() string {
	return "mailto:" + c.Email
}

func (c Content) WiFiString() string {
	hidden := "false"
	if c.IsHidden {
		hidden = "true"
	}
	// TODO: validate
	typ := c.Type
	return fmt.Sprintf("WIFI:S:%s;T:%s;P:%s;H:%s;", c.SSID, typ, c.Password, hidden)
}

func (c Content) TelString() string {
	panic("TODO: not implemented")
	return ""
}

func (c Content) ContactString() string {
	panic("TODO: not implemented")
	return ""
}
