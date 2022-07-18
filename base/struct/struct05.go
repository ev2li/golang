package main

import (
	"fmt"
	"unsafe"
)

type Person02 struct {
	name string
	sex  byte
	age  int
}

func main() {
	var man4 Person02 = Person02{"andy", 'm', 20}
	fmt.Printf("&man4 = %p\n", &man4)
	fmt.Printf("&man4.name = %p\n", &(man4.name))
	fmt.Println("man4 size:", unsafe.Sizeof(man4))

	var man6 *Person02 = &Person02{"zhangli", 'm', 31}
	fmt.Println("man6:", man6)
	var man7 *Person02 = new(Person02)
	fmt.Println("man7:", man7)
}
