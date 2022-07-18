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

	man2 := Person{name: "Rose", age: 30}
	fmt.Println("man2:", man2)
	//索引成员变量
	fmt.Println("man2.name:", man2.name)

	var man3 Person
	man3.name = "Mike"
	man3.sex = 'm'
	man3.age = 40
	fmt.Println("man3:", man3)

	var tmp Person
	fmt.Println("tmp:", tmp)
	tmp = man3
	fmt.Println("tmp:", tmp)
}
