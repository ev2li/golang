package main

import (
	"fmt"
	"strings"
)

func wordCountFunc(str string) map[string]int {
	srt := strings.Fields(str) //将符串拆分成字符串切片
	rt := make(map[string]int)

	for _, word := range srt {
		if v, ok := rt[word]; ok {
			//存在
			rt[word] = v + 1
		} else {
			rt[word] = 1
		}
	}
	return rt
}

func main() {
	str := "I lov my work and I love my family too"
	mrt := wordCountFunc(str)
	fmt.Println("mrt:", mrt)
}
