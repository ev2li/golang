package main

import (
	"fmt"
	"unsafe"
)

type Person01 struct {
	name string
	sex  byte
	age  int
}

func test1(man *Person01) {
	man.name = "zhangli"
	man.sex = 'm'
	man.age = 31
	fmt.Println("man:", man)
	fmt.Println("sizeof(man):", unsafe.Sizeof(man))
}

func main() {
	var tmp Person01
	test1(&tmp) //指针传递
	fmt.Println("tmp:", tmp)
}
