package main

import "fmt"

func noEmpty(data []string) []string {
	out := data[:0] //在原切片上截取一个长度为0的切片
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}

	return out
}

func main() {
	//1.自动推导赋初始值
	data := []string{"red", "", "black", "", "", "pink", "blue"}
	afterData := noEmpty(data)
	fmt.Println("afterData:", afterData)
}
