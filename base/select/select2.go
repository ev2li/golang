package main

import (
	"fmt"
	"time"
)

//select超时处理

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() { //子go程获取数据
		for {
			select {
			case num := <-ch:
				fmt.Println("num=", num)
			case <-time.After(time.Second * 3):
				quit <- true
				goto lable
			}
		}
	lable:
		fmt.Println("break to lable")
	}()
	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}
	<-quit
	fmt.Println("finish!!!")
}
