package main

import "fmt"

type Person03 struct {
	name      string
	age       int
	flag      bool
	intereset []string
}

//通过函数参数初始化结构体
func initFunc(p *Person03) {
	p.name = "andy"
	p.age = 20
	p.flag = true
	p.intereset = append(p.intereset, "yuwen")
	p.intereset = append(p.intereset, "shopping")
}

//通过函数返回值来初始化结构体

func initFunc2() *Person03 {
	p := new(Person03)
	p.name = "andy"
	p.age = 20
	p.flag = true
	p.intereset = append(p.intereset, "yuwen")
	p.intereset = append(p.intereset, "shopping")
	return p
}

func main() {
	var man4 Person03
	initFunc(&man4)
	fmt.Println("man4:", man4)

	man5 := initFunc2()
	fmt.Println("man5:", man5)
}
