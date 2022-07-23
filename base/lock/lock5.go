package main

/*
	go语言中互斥锁，读写锁不要和channel混用---隐性死锁
*/
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg02 sync.WaitGroup
var rwmutex02 sync.RWMutex

func readGo02(in chan int, n int) {
	defer wg02.Done()
	for {
		num := <-in
		fmt.Printf("%d读go程,读取%d\n", n, num)
	}
}

func writeGo02(out chan int, n int) {
	defer wg02.Done()
	for {
		//生成随机数
		num := rand.Intn(1000)
		rwmutex02.Lock()
		out <- num
		fmt.Printf("%d写go程,写入%d\n", n, num)
		rwmutex02.Unlock()
		time.Sleep(time.Microsecond)
	}
}

func main() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())
	wg02.Add(10)
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go readGo02(ch, i) //5个读go程
	}

	for i := 0; i < 5; i++ {
		go writeGo02(ch, i) //5个写go程
	}

	wg02.Wait()
}
