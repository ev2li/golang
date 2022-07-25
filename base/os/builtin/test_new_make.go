package main

import "fmt"

func testNew() {
	b := new(bool) //返回的是一个指针
	fmt.Printf("b: %T\n", b)
	fmt.Printf("b: %v\n", *b)
	i := new(int)
	fmt.Printf("i: %T\n", i)
	fmt.Printf("i: %v\n", *i)

	s := new(string)
	fmt.Printf("s: %T\n", s)
	fmt.Printf("s: %v\n", *s)
}

func testMake() {
	var p *[]int = new([]int)
	fmt.Printf("p: %v\n", p)
	v := make([]int, 10)
	fmt.Printf("v: %v\n", v)
}

func main() {
	// testNew()
	testMake()
}
