package main

import (
	"fmt"
	"time"
)

func test01() {
	t := time.Now()
	fmt.Printf("t: %T\n", t)
	fmt.Printf("t: %v\n", t)
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T,%T,%T,%T,%T,%T\n", year, month, day, hour, minute, second)

}

func test_timestamp() {
	t := time.Now()
	fmt.Printf("t.Unix(): %T\n", t.Unix())
	fmt.Printf("t.Unix(): %v\n", t.Unix())
}

func test_timeStampNano() {
	t := time.Now()
	fmt.Printf("t.UnixNano(): %T\n", t.UnixNano())
	fmt.Printf("t.UnixNano(): %v\n", t.UnixNano())
}

func timeStampToTime() {
	timestamp := time.Now().Unix()
	t := time.Unix(timestamp, 0)
	fmt.Println(t)
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func main() {
	// test01()
	// test_timestamp()
	// test_timeStampNano()
	timeStampToTime()
}
