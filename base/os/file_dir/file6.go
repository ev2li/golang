package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f_s, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer f_s.Close()
	fmt.Println("success")

	f_d, err := os.Create("./d.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer f_d.Close()

	//从读文件中读取数据,放到缓冲区
	buf := make([]byte, 4*1024)
	//循环从读文件中获取数据,原封不动的写到文件里
	for {
		n, err := f_s.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读完:", n)
			return
		}
		f_d.Write(buf[:n]) //读多少,写多少
	}
}
