package main

import (
	"fmt"
	"runtime"
)

type paincContext struct {
	function string
}

func ProtectRun(entry func()) {
	//延迟处理函数
	defer func() {
		//发生宕机时，获取painc传递上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("runtime error:", err)
		default: //非运行时错误
			fmt.Println("error:", err)
		}

	}()
	entry()
}

func main() {
	fmt.Println("运行前")
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		panic(&paincContext{
			"手动触发painc",
		})
		fmt.Println("手动宕机后")
	})

	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")

	})
	fmt.Println("运行后")
}
