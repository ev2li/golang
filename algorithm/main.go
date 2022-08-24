package main

import (
	"algorithm/ArrayList"
	"algorithm/StackArray"
	"fmt"
)

func main1() {
	myStatck := StackArray.NewStack()
	myStatck.Push(1)
	myStatck.Push(2)
	myStatck.Push(3)
	myStatck.Push(4)

	fmt.Println(myStatck.Pop())
	fmt.Println(myStatck.Pop())
	fmt.Println(myStatck.Pop())
	fmt.Println(myStatck.Pop())
}

func main2() {
	myStatck := ArrayList.NewArrayListStack()
	myStatck.Push(1)
	myStatck.Push(2)
	myStatck.Push(3)
	myStatck.Push(4)

	fmt.Println(myStatck.Pop())
	fmt.Println(myStatck.Pop())
	fmt.Println(myStatck.Pop())
	fmt.Println(myStatck.Pop())
}

func main() {
	myStatck := ArrayList.NewArrayListStackIterator()
	myStatck.Push(1)
	myStatck.Push(2)
	myStatck.Push(3)
	myStatck.Push(4)

	for it := myStatck.Iterator; it.HasNext(); {
		item, _ := it.Next()
		fmt.Println(item)
	}
}
