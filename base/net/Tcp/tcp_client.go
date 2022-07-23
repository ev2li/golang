package main

import (
	"fmt"
	"net"
)

func main() {
	//创建连接
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial() err:", err)
		return
	}
	defer conn.Close()
	buf := []byte("Are you Ready")
	conn.Write(buf)

	//接收服务器回发的数据
	buf = make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("服务器回发:", string(buf[:n]))

}
