package main

/*
channel实现生产者消费者模型
*/

import "fmt"

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
		fmt.Println("生产:", i*i)
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("消费者拿到:", num)
	}
}

func main() {
	ch := make(chan int, 5)
	go producer(ch) //子go程生产者
	consumer(ch)    //主go程是消费者
}
