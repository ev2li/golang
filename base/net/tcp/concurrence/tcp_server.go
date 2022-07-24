package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}

		//具体完成服务器与客户端的数据通信
		go func(conn net.Conn) {
			defer conn.Close()
			addr := conn.RemoteAddr()
			fmt.Println(addr, "客户端成功连接")
			for {
				buf := make([]byte, 4096)
				n, err := conn.Read(buf)
				if string(buf[:n]) == "exit\n" {
					fmt.Println("服务器接收到客户退出请求，断开连接!!")
					return
				}

				if n == 0 {
					fmt.Println("服务器检测到客户端已关闭，断开连接!!!")
					return
				}

				if err != nil {
					fmt.Println("conn.Read err:", err)
					return
				}
				fmt.Println("服务端读取了:", string(buf[:n]))
				//小写转大写
				conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
			}
		}(conn)
	}
}
