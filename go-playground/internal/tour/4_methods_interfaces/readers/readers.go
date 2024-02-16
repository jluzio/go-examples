package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/reader"
)

func main() {
	// main_reader()
	// exercise_reader()
	exercise_rot13Reader()
}

func main_reader() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(buffer []byte) (int, error) {
	for i := range buffer {
		buffer[i] = 'A'
	}
	return len(buffer), nil
}
func exercise_reader() {
	reader.Validate(MyReader{})
}

type rot13Reader struct {
	r io.Reader
}

func rot13(v byte) byte {
	var anchorChar byte

	if v >= 'a' && v <= 'z' {
		anchorChar = 'a'
	} else if v >= 'A' && v <= 'Z' {
		anchorChar = 'A'
	}

	if anchorChar == 0 {
		return v
	}

	letterIndex := v - anchorChar
	updatedLetterIndex := (letterIndex + 13) % 26
	updatedByte := updatedLetterIndex + anchorChar
	return updatedByte
}

func (rot rot13Reader) Read(buffer []byte) (int, error) {
	n_read, err := rot.r.Read(buffer)
	if err == nil {
		for i, v := range buffer {
			buffer[i] = rot13(v)
		}
	}
	return n_read, err
}

func exercise_rot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
