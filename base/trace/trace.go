package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

//trace编程过程
//创建文件
//启动
//停止
func main() {
	//1.创建一个trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//2.启动treace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	//正常要调度的业务
	fmt.Println("Hello GMP")
	//3.停止
	trace.Stop()
}
