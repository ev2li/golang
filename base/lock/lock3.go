package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var rwmutex sync.RWMutex

func readGo(in <-chan int, n int) {
	defer wg1.Done()
	for {
		rwmutex.RLock()
		num := <-in
		fmt.Printf("%d读go程,读取%d\n", n, num)
		rwmutex.RUnlock()
	}
}

func writeGo(out chan<- int, n int) {
	defer wg1.Done()
	for {
		//生成随机数
		num := rand.Intn(1000)
		rwmutex.Lock()
		out <- num
		fmt.Printf("%d写go程,写入%d\n", n, num)
		rwmutex.Unlock()
		time.Sleep(time.Microsecond)
	}
}

func main() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())
	wg1.Add(10)
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go readGo(ch, i)
	}

	for i := 0; i < 5; i++ {
		go writeGo(ch, i)
	}

	wg1.Wait()
}
