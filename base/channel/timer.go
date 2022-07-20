package main

import (
	"fmt"
	"time"
)

/*
定时器
*/

func main() {
	/*fmt.Println(time.Now()) //当前时间
	//创建一个定时器
	myTimer := time.NewTicker(time.Second * 2)
	nowTime := <-myTimer.C
	fmt.Println(nowTime)*/
	//1.sleep
	time.Sleep(time.Second)
	//2. Timer.C
	myTimer := time.NewTicker(time.Second * 2)
	nowTime := <-myTimer.C
	fmt.Println(nowTime)
	//3.time.After()
	nowTime = <-time.After(time.Second * 2)

}
