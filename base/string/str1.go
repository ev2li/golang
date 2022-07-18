package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I love my work and I love my family too"
	arr := strings.Split(str, "I") //字符串按指定分隔符拆分
	fmt.Println("arr:", arr)

	//字符串按空格拆分
	arr2 := strings.Fields(str)
	fmt.Println("arr2:", arr2)

	//判断字符串结束标记
	flag := strings.HasSuffix("test.abc", ".abc")
	fmt.Println("flag:", flag)

	flag = strings.HasPrefix("xx.test", "xy")
	fmt.Println("flag:", flag)
}
