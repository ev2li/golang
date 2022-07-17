package main

import "fmt"

func noEmpty2(data []string) []string {
	i := 0
	for _, str := range data {
		if str != "" {
			data[i] = str
			i++
		}
	}
	return data[0:i]
}

func main() {
	//1.自动推导赋初始值
	data := []string{"red", "", "black", "", "", "pink", "blue"}
	afterData := noEmpty2(data)
	fmt.Println("afterData:", afterData)
}
