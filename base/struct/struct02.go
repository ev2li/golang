package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func main() {
	//1.顺序初始化
	var man Person = Person{"andy", 'm', 20}
	fmt.Println("man:", man)

	var man3 Person
	man3.name = "Mike"
	man3.sex = 'm'
	man3.age = 40
	fmt.Println("man3:", man3)

	fmt.Println("man和man比较:", man == man3)
}
