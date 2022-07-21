package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var mutex01 sync.Mutex //创建一个互斥锁

func printer1(str string) {
	mutex01.Lock()
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex01.Unlock()
}

func person01() {
	defer wg.Done()
	printer1("hello")
}

func person02() {
	defer wg.Done()
	printer1("world")
}

func main() {
	wg.Add(2)
	go person01()
	go person02()
	wg.Wait()
}
