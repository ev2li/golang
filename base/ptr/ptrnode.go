package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println("&a", &a)
	var p *int = &a
	fmt.Println(*p)

	//在heap上申请一片内存地址空间
	//p = new(int) //默认类型的，默认值

	var n *string
	n = new(string)
	fmt.Printf("%q\n", *n)
}
