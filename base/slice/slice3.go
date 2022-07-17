package main

import "fmt"

func main() {
	//1.自动推导赋初始值
	s1 := []int{1, 2, 3, 4}
	s1 = append(s1, 888)
	s1 = append(s1, 888)
	s1 = append(s1, 888)
	s1 = append(s1, 888)
	fmt.Println("s1=", s1)
}
