package main

import (
	"fmt"
	"time"
)

func main() {
	/* 	timer := time.NewTimer(time.Second * 2)
	   	fmt.Printf("time.Now(): %v\n", time.Now())
	   	t1 := <-timer.C //阻塞的，直到指定时间到了
	   	fmt.Printf("t1: %v\n", t1) */
	/* fmt.Printf("time.Now(): %v\n", time.Now())
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Printf("time.Now(): %v\n", time.Now())
	*/

	// time.Sleep(time.Second * 2)
	// fmt.Printf("time.Now(): %v\n", time.Now())
	// <-time.After(time.Second * 2)
	// fmt.Printf("time.Now(): %v\n", time.Now())

	/* 	timer := time.NewTicker(time.Second)
	   	go func() {
	   		<-timer.C
	   		fmt.Println("func...")
	   	}()
	   	timer.Stop()
	   	time.Sleep(time.Second * 3)
	   	fmt.Println("main end...") */
	fmt.Println("before...")
	t := time.NewTimer(time.Second * 5)
	t.Reset(time.Second * 1)
	<-t.C
	fmt.Println("after")
}
