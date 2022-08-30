package main

import (
	"algorithm/Link"
	"fmt"
)

func main1() {
	node1 := new(Link.Node)
	node2 := new(Link.Node)
	node3 := new(Link.Node)
	node4 := new(Link.Node)
	node1.Data = 1
	node1.PNext = node2
	node2.Data = 2
	node2.PNext = node3
	node3.Data = 3
	node3.PNext = node4
	node4.Data = 4
	fmt.Println(node1.Data)
	fmt.Println(node2.Data)
	fmt.Println(node3.Data)
	fmt.Println(node4.Data)

	fmt.Printf("node1.PNext.Data: %v\n", node1.PNext.PNext.Data)
	fmt.Printf("node1.Length(): %v\n", node1.Length())
}

func main2() {
	mystack := Link.NewStack()
	for i := 0; i < 1000; i++ {
		mystack.Push(i)
	}

	for data, _ := mystack.Pop(); data != nil; data, _ = mystack.Pop() {
		fmt.Printf("data: %v\n", data)
	}
}

func main() {
	myq := Link.NewLinkQueue()
	for i := 0; i < 100; i++ {
		myq.EnQueue(i)
	}

	for i := 0; i < 100; i++ {
		value := myq.DeQueue()
		fmt.Println(value)
	}

}
