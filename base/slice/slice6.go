package main

import "fmt"

func noRepeat(data []string) []string {
	rt := data[:1]
	flag := false

	for _, str := range data {
		for _, v := range rt {
			if str == v {
				flag = true
				break
			}
			flag = false
		}
		if !flag {
			rt = append(rt, str)
		}
	}
	return rt
}

func main() {
	//1.自动推导赋初始值
	data := []string{"red", "black", "black", "blue", "red", "pink", "blue"}
	afterData := noRepeat(data)
	fmt.Println("afterData:", afterData)
}
