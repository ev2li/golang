package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats-streaming-server/logger"
)

func test_log() {
	log.Print("my log")
	log.Printf("my log %d", 100)
	name := "tom"
	age := 20

	log.Println(name, "+", age)
}

func test_log_02() {
	defer fmt.Println("painc结束后再执行...")
	log.Print("my log")
	log.Panic("my painc")
	fmt.Println("end...")
}
func test_log_03() {
	defer fmt.Println("fatal结束后再执行...")
	log.Print("my log")
	log.Fatal("my Fatal")
	os.Exit(1)
	fmt.Println("end...")
}

/* func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetPrefix("my log")
	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal("日志文件错误")
	}
	log.SetOutput(f)
} */

func test_log_04() {

	i := log.Flags()
	fmt.Printf("i: %v\n", i)
	log.Print("my log...")
}

func init() {
	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal("日志文件错误")
	}

	logger := log.New(f, "MyLog: ", log.Ldate|log.Ltime|log.Llongfile)
}

func main() {
	logger.Print("my log...")
}
