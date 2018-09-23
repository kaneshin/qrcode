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
	fi, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	var r io.Reader
	if fi.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		r = os.Stdin
	} else {
		args := flag.Args()
		if len(args) == 0 {
			r = os.Stdin
		} else {
			f, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer f.Close()
			r = f
		}
	}

	return do(r)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
