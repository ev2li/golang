package main

import (
	"fmt"
	"os"
)

func main() {
	if f, err := os.Open("./2.txt"); err != nil { //以只读方式打开
		fmt.Println("打开文件出错")
	} else {
		fmt.Println("success")
		_, err := f.WriteString("#####")
		if err != nil {
			fmt.Println("写入错误")
		}
		defer f.Close()
	}

}
