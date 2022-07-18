package main

import (
	"fmt"
	"os"
)

func main() {
	if f, err := os.Create("./1.txt"); err != nil {
		fmt.Println("打开文件出错")
	} else {
		fmt.Println("success")
		defer f.Close()
	}
}
