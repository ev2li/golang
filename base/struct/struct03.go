package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func test(man Person) {
	man.name = "zhangli"
	man.sex = 'm'
	man.age = 31
	fmt.Println("man:", man)
}

func main() {
	var tmp Person
	test(tmp) //值传递
	fmt.Println("tmp:", tmp)
}
