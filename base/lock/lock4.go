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

var wg01 sync.WaitGroup
var rwmutex01 sync.RWMutex
var value int //定义全局变量，模拟共享数据

func readGo01(n int) {
	defer wg01.Done()
	for {
		rwmutex01.RLock()
		num := value
		fmt.Printf("%d读go程,读取%d\n", n, num)
		rwmutex01.RUnlock()
	}
}

func writeGo01(n int) {
	defer wg01.Done()
	for {
		//生成随机数
		num := rand.Intn(1000)
		rwmutex01.Lock()
		value = num
		fmt.Printf("%d写go程,写入%d\n", n, num)
		rwmutex01.Unlock()
		time.Sleep(time.Microsecond)
	}
}

func main() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())
	wg01.Add(10)

	for i := 0; i < 5; i++ {
		go readGo01(i) //5个读go程
	}

	for i := 0; i < 5; i++ {
		go writeGo01(i) //5个写go程
	}

	wg01.Wait()
}
