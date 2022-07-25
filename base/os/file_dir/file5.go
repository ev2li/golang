package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("./1.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("文件打开失败")
	}
	defer f.Close()
	fmt.Println("success")

	reader := bufio.NewReader(f)
	buf, err := reader.ReadBytes('\n') //读一行
	if err == io.EOF {
		fmt.Println("文件读取完毕")
	} else if err != nil {
		fmt.Println("ReadBytes failed:", err)
	}

	fmt.Println(string(buf))

}
