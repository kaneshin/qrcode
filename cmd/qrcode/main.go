package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kaneshin/qrcode"
)

var (
	decode = flag.Bool("d", false, "decode data")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: qrcode [OPTION]... [FILE]
QRcode encode or decode FILE, or standard input, to standard output.
`)
		flag.PrintDefaults()
	}
	flag.Parse()
}

func doEncode(dst io.Writer, src []byte) error {
	enc := qrcode.NewEncoder(dst)
	return enc.Encode(src)
}

func doDecode(dst io.Writer, src []byte) error {
	buf := bytes.NewBuffer(src)
	dec := qrcode.NewDecoder(buf)
	return dec.Decode(dst)
}

func do(r io.Reader) error {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if *decode {
		err = doDecode(&buf, src)
	} else {
		err = doEncode(&buf, src)
	}
	if err != nil {
		return err
	}

	os.Stdout.Write(buf.Bytes())
	return nil
}

func run() error {
	var name string
	if args := flag.Args(); len(args) > 0 {
		name = args[0]
	}

	var r io.Reader
	switch name {
	case "", "-":
		r = os.Stdin
	default:
		f, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		r = f
	}

	return do(r)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
