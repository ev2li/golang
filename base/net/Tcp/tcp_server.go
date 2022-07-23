package main

import (
	"fmt"
	"net"
)

func main() {
	//指定服务器，地址和端口 创建一个用于监听的套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("net listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务器等待客户建立连接")

	//阻塞监听客户端连接请求
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept() err:", err)
		return
	}
	defer conn.Close()

	fmt.Println("服务器与客户端成功建立连接")
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read() err:", err)
		return
	}
	conn.Write(buf[:n])
	//处理数据
	fmt.Println("服务器读到数据:", string(buf[:n]))
}
