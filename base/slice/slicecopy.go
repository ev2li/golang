package main

import "fmt"

func main() {
	//1.自动推导赋初始值
	s1 := []int{0, 1, 2, 3, 4}
	s2 := []int{7, 8, 9}

	copy(s1, s2)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}
