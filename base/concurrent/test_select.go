package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {
	go func() {
		chanInt <- 100
		chanStr <- "hello"
		defer close(chanInt)
		defer close(chanStr)
	}()

	for {
		select {
		case result := <-chanInt:
			fmt.Printf("result: %v\n", result)
		case result := <-chanStr:
			fmt.Printf("result: %v\n", result)
		default:
			fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
}
