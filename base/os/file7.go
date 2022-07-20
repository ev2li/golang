package main

import (
	"fmt"
	"os"
)

/*
	目录操作
*/
func main() {
	fmt.Println("请输入待查询的目录")
	var path string
	fmt.Scan(&path)
	fptr, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err:", err)
	}
	defer fptr.Close()

	//读取目录项
	infos, err := fptr.Readdir(-1)
	if err != nil {
		fmt.Println("Readdir err:", err)
	}

	for _, info := range infos {
		if info.IsDir() {
			fmt.Println(info.Name(), "是一个目录")
		} else {
			fmt.Println(info.Name(), "是一个文件")
		}

	}
}
