package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("-------I'm goroutine--------")
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("-------I'm main--------")
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}

}
