package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(1) //设置单核cpu
	fmt.Println("n = ", n)
	n = runtime.GOMAXPROCS(2) //设置双核cpu
	fmt.Println("n = ", n)
	fmt.Println(runtime.GOROOT())
	/*for {
		go fmt.Print(0) //子go程
		fmt.Print(1)    //主go程
	}*/

}
