package main

import (
	"flag"
	"fmt"

	"github.com/kaneshin/qrcode"
)

func main() {
	c := qrcode.Content{}
	flag.StringVar(&c.URL, "url", "", "")
	flag.StringVar(&c.SSID, "ssid", "", "")
	flag.StringVar(&c.Password, "password", "", "")
	flag.StringVar(&c.Type, "type", "", "")
	flag.Parse()
	fmt.Print(c.String())
}
