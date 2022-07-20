package main

import "fmt"

func main() {
	ch := make(chan string) //无缓冲channel
	fmt.Println("len(ch):", len(ch), "- cap(ch):", cap(ch))
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("i = ", i)
		}
		//通知主go程打印完毕
		ch <- "子go程打印完毕"
	}()
	str := <-ch
	fmt.Println("str:", str)
}
