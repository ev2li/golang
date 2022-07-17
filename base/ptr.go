package base2

import "fmt"

func test(m int) {
	var b int = 1000
	b += m
	fmt.Println("b:", b)
}

func main() {
	var a int = 10
	var p *int = &a

	a = 100
	fmt.Println("a = ", a)
	*p = 250
	fmt.Println("a = ", a)
	a = 1000
	fmt.Println("*p = ", *p)
	test(a)
}
