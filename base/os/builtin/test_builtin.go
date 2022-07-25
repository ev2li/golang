package main

import "fmt"

func testAppend() {
	s := []int{1, 2, 3}
	s = append(s, 100)
	fmt.Printf("s: %v\n", s)
	s1 := []int{4, 5, 6}
	s1 = append(s, s1...)
	fmt.Printf("s1: %v\n", s1)
	name := []byte("aaaaa")
	fmt.Printf("string(name): %v\n", string(name))
}
func testLen() {
	s := "Hello World"
	fmt.Printf("len(s): %v\n", len(s))
	s1 := []int{1, 2, 3}
	fmt.Printf("len(s1): %v\n", len(s1))
}

func main() {
	defer fmt.Println("panic后,我还会执行...")
	panic("完蛋了...")
	fmt.Println("end...")
}
