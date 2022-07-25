package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func test_json() {
	p := Person{
		Name:  "tom",
		Age:   30,
		Email: "123@qq.com",
	}
	b, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("string(b): %v\n", string(b))
	}

	var p2 Person
	if err := json.Unmarshal(b, &p2); err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("p2: %v\n", p2)
	}
	fmt.Printf("p2.Name: %v\n", p2.Name)
}

func main() {
	test_json()
}
