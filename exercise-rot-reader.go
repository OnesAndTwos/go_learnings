package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (num int, err error) {

	num, err = reader.r.Read(b)

	for i := 0; i < num; i++ {

		if b[i] > 64 && b[i] < 91 {
			b[i] = (b[i]+13)%65 + 65
		} else if b[i] > 96 && b[i] < 123 {
			b[i] = (b[i]+13)%97 + 97
		}

	}

	return num, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
