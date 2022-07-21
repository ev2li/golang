package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)    //用来进行数据通信的channel
	quit := make(chan bool) //用来判断是否退出的channel
	// ch2 := make(chan string)

	go func() { //子go程写数据
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true //通知主go程退出
		runtime.Goexit()
	}()

	for { //主go程读数据
		select {
		case num := <-ch:
			fmt.Println("num:", num)
		case <-quit:
			// break //break跳出select
			// runtime.Goexit() 不能退出主协程
			return
		}
	}
}
