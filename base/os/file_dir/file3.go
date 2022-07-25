package main

import (
	"fmt"
	"os"
)

func main() {
	if f, err := os.OpenFile("./1.txt", os.O_RDWR, 6); err != nil { //必须是存在的文件
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
