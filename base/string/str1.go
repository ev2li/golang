package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I love my work and I love my family too"
	arr := strings.Split(str, " ")
	fmt.Println("arr:", arr)
}
