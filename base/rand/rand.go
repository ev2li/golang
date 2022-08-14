package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	min := 10
	max := 20
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(50))            //生成0-50间的数
	fmt.Println(rand.Intn(max-min) + min) //生成10-20间的数

	fmt.Println(rand.Float64()) //0.0-1.0间的数

}
