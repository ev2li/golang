package main

import (
	"fmt"
	"os"
)

func main() {
	//获取所有环境变量
	s := os.Environ()
	fmt.Printf("s: %v\n", s)
	// 获取某个环境变量
	s2 := os.Getenv("GOPATH")
	fmt.Printf("s2: %v\n", s2)

	//设置环境变量
	os.Setenv("env1", "env1")
	s3 := os.Getenv("env1")
	fmt.Printf("s3: %v\n", s3)
	fmt.Println("----------------")

	// 查找
	s4, b := os.LookupEnv("env1")
	fmt.Printf("b: %v\n", b)
	fmt.Printf("s4: %v\n", s4)

	//替换
	os.Setenv("NAME", "gohper")
	os.Setenv("BURROW", "/usr/gopher")
	fmt.Printf("os.ExpandEnv(\" lives in \"): %v\n", os.ExpandEnv(" lives in "))

	//清空环境变量(不要操作，知道就行)
	// os.Clearenv()
}
