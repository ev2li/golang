package main

import "fmt"

func remove(data []int, index int) []int {
	copy(data[index:], data[index+1:])
	return data[:len(data)-1]
}

func main() {
	//1.自动推导赋初始值
	s1 := []int{5, 6, 7, 8, 9}

	afterData := remove(s1, 2)
	fmt.Println("afterData:", afterData)
}
