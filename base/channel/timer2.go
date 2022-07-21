package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)
	fmt.Println("now:", time.Now())
	myTicker := time.NewTimer(time.Second)
	i := 0
	go func() {
		for {
			nowTime := <-myTicker.C
			i++
			fmt.Println("nowTime:", nowTime)
			for i == 8 {
				quit <- true
			}
		}
	}()

	<-quit
}
