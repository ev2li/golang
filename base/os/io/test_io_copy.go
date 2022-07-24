package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func testCopy() {
	r := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)
	//buf的copy效率更高
	if _, err := io.CopyBuffer(os.Stdout, r, buf); err != nil {
		log.Fatal(err)
	}

	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}

func testIO() {
	pr, pw := io.Pipe()
	go func() {
		fmt.Fprint(pw, "some io.Reader stream to be read\n")
		pw.Close()
	}()

	if _, err := io.Copy(os.Stdout, pr); err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	testCopy()

	testIO()
}
