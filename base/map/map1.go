package main

import "fmt"

func main() {
	var m1 map[int]string //map声明，没有空间，不能直接存储key -- value
	// m1[100] = "Green"
	if m1 == nil {
		fmt.Println("map is nil")
	}

	m2 := map[int]string{} //
	fmt.Println("len(m2):", len(m2))
	fmt.Println("m2:", m2)
	m2[4] = "red"
	fmt.Println("m2:", m2)

	m3 := make(map[int]string)
	fmt.Println("len(m3):", len(m3))
	fmt.Println("m3:", m3)
	m3[2] = "pink"
	fmt.Println("m3:", m3)

	m4 := make(map[int]string, 5)
	fmt.Println("len(m4):", len(m4))
	fmt.Println("m4:", m4)
	m4[2] = "pink"
	fmt.Println("m4:", m4)

	//初始化map
	var m5 map[int]string = map[int]string{1: "red", 2: "green"}
	fmt.Println("m5:", m5)
}
