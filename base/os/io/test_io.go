package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("Hello World")
	buf := make([]byte, 20)
	r.Read(buf)
	fmt.Printf("string(buf): %v\n", string(buf))
}
