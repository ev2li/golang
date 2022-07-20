package main

import (
	"fmt"
	"runtime"
)

/*func test() {
	fmt.Println("ccccccc")
	return
	defer fmt.Println("dddddddd")
}*/

func test() {
	defer fmt.Println("ccccccc")
	runtime.Goexit() //退出当前go程
	fmt.Println("dddddddd")
}

func main() {
	go func() {
		fmt.Println("aaaaa")
		test()
		defer fmt.Println("bbbbb")
	}()
	for {

	}
}
