package main

import "fmt"

// OCP设计原则
type Pet02 interface {
	eat02()
	sleep02()
}

type Dog02 struct {
}

type Cat02 struct {
}

func (dog02 Dog02) eat02() {
	fmt.Println("dog eat...")
}

func (dog02 Dog02) sleep02() {
	fmt.Println("dog sleep...")
}

func (cat02 Cat02) eat02() {
	fmt.Println("cat eat...")
}

func (cat02 Cat02) sleep02() {
	fmt.Println("cat sleep...")
}

type Person struct {
}

//pet即可以传递Dog02也可以传递Cat02
func (p Person) care(pet Pet02) {
	pet.eat02()
	pet.sleep02()
}

func main() {
	dog02 := Dog02{}
	cat02 := Cat02{}

	person := Person{}
	person.care(dog02)
	person.care(cat02)
}
