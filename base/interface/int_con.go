package main

import "fmt"

/*模拟构造函数*/

type Person03 struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person03, error) {
	if name == "" {
		return nil, fmt.Errorf("name不能为空")
	}

	if age < 0 {
		return nil, fmt.Errorf("age不能小于0")
	}
	return &Person03{name: name, age: age}, nil
}

func main() {
	p, err := NewPerson("Tom", 20)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("p: %v\n", p)
	}
}
