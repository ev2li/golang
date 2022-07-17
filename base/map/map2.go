package main

import "fmt"

func main() {
	m7 := make(map[int]string, 1)
	m7[100] = "Nami"
	m7[20] = "Hello"
	m7[3] = "World"

	fmt.Println("m5:", m7)

	m7[3] = "Yellow"
	fmt.Println("m5:", m7) //将原map中key值为3的map元素覆盖
	for k, v := range m7 {
		fmt.Println("key-value:", k, v)
	}

	//判断map中key是否存在
	if v, ok := m7[20]; ok {
		fmt.Println("存在v:", v)
	}
}
