package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex //创建一个互斥锁

func printer(str string) {
	mutex.Lock() //先加锁,再访问
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock()
}

func person1() {
	printer("hello")
}

func person2() {
	printer("world")
}

func main() {
	go person1()
	go person2()
	for { //ugly

	}
}
