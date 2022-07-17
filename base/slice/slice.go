package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	/*s := arr[1:3:5]

	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s))
	fmt.Println("cap(s)=", cap(s))*/

	/*s := arr[1:5:7]
	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s)) //5-1 == 4
	fmt.Println("cap(s)=", cap(s)) //7-1 == 6

	s2 := s[:6]
	fmt.Println("s2=", s2)
	fmt.Println("len(s2)=", len(s2)) // 6 - 0 == 6
	fmt.Println("cap(s2)=", cap(s2)) // 7 - 1 == 6*/

	s := arr[2:5] //{3, 4, 5}
	fmt.Println("s:", s)
	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s)) //5-2 == 3
	fmt.Println("cap(s)=", cap(s)) //10-2 == 8

	s2 := arr[2:7]
	// s2 := s[2:7] painc out of range
	fmt.Println("s2:", s2) //3 4 5 6 7
	fmt.Println("s2=", s2)
	fmt.Println("len(s2)=", len(s2)) // 7 - 2 == 5
	fmt.Println("cap(s2)=", cap(s2)) // 没有指定那就是 10 - 2 == 8
}
