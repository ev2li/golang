package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("---正在唱:月亮之上---")
		time.Sleep(100 * time.Millisecond)
	}

}

func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("---正在跳: 赵四---")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sing()
	go dance()
	for {

	}
}
