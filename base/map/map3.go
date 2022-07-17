package main

import "fmt"

func main() {
	m := make(map[int]string, 1)
	m[100] = "Nami"
	m[20] = "Hello"
	m[3] = "World"

	delete(m, 3)
	fmt.Println("m:", m)
}
