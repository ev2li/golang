package main

import "fmt"

func main() {
	ch := make(chan int) //无缓冲channel
	fmt.Println("len(ch):", len(ch), "- cap(ch):", cap(ch))
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子go程 ", i)
			ch <- i
		}
	}()
	for i := 0; i < 5; i++ {
		n := <-ch
		fmt.Println("主go程", n)
	}
}
