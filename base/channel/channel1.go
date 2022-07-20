package main

import (
	"fmt"
	"time"
)

//全局定义channel,用来完成通信
var channel = make(chan int)

//定义一个打印机
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(time.Second)
	}
}

//定义两个人使用打印机
func p1() { //p1先执行
	printer("Hello")
	channel <- 8
}

func p2() {
	<-channel
	printer("world")
}

func main() {
	go p1()
	go p2()
	for {

	}
}
