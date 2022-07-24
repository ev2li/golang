package main

import (
	"fmt"
	"sync/atomic"
)

func test_add_sub() {
	var i int32 = 100
	atomic.AddInt32(&i, 1) //增减
	fmt.Printf("i: %v\n", i)
}
func test_load_store() {
	var i int32 = 100
	atomic.LoadInt32(&i)
	fmt.Printf("i: %v\n", i)   //载入read
	atomic.StoreInt32(&i, 200) //载入write
	fmt.Printf("i: %v\n", i)
}

func test_cas() {
	//cas
	var i int32 = 100
	b := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("i: %v\n", i)
}

func main() {
	test_cas()
}
