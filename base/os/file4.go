package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("./1.txt", os.O_RDWR, 6)
	if err != nil { //必须是存在的文件
		fmt.Println("打开文件出错")
	}
	defer f.Close()

	fmt.Println("success")
	n, err := f.WriteString("#####")
	if err != nil {
		fmt.Println("写入错误")
	}
	fmt.Println("写了n个字符:", n)
	f.Seek(6, io.SeekStart)
	f.WriteString("test")
	offset, err := f.Seek(-5, io.SeekEnd)
	fmt.Println("offset:", offset)

	n, _ = f.WriteAt([]byte("11111111111"), offset)
	fmt.Println("n:", n)
}
