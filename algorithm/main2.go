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

func main3() {
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

func Add(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + Add(num-1)
	}
}

func main4() {
	// fmt.Printf("Add(5): %v\n", Add(5))
	mystatck := StackArray.NewStack()
	mystatck.Push(5)
	last := 0 //保存结果
	for !mystatck.IsEmpty() {
		data := mystatck.Pop() //取出数据

		if data == 0 {
			last += 0
		} else {
			last += data.(int)
			mystatck.Push((data.(int) - 1))
		}
	}
	fmt.Printf("last: %v\n", last)
}

func fab(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return fab(num-1) + fab(num-2)
	}
}

func main() {
	fmt.Printf("fab(7): %v\n", fab(7))

	mystatck := StackArray.NewStack()
	mystatck.Push(7)
	last := 0 //保存结果
	for !mystatck.IsEmpty() {
		data := mystatck.Pop() //取出数据

		if data == 1 || data == 2 {
			last += 1
		} else {
			mystatck.Push((data.(int) - 2))
			mystatck.Push((data.(int) - 1))
		}
	}
	fmt.Printf("last: %v\n", last)
}
