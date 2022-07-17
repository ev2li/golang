package main

import "fmt"

func main() {
	//1.自动推导赋初始值
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1=", s1)

	s2 := make([]int, 5, 10)
	fmt.Println("len(s2)=", len(s2))
	fmt.Println("cap(s2)=", cap(s2))

	s3 := make([]int, 10)
	fmt.Println("len(s3)=", len(s3))
	fmt.Println("cap(s3)=", cap(s3))
}
