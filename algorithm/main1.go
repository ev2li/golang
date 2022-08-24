package main

import (
	"algorithm/ArrayList"
	"fmt"
)

func main1() {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Println(list)
}

func main2() {
	list := ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list)
	fmt.Printf("list.TheSize: %v\n", list.TheSize)
}

func main4() {
	var list *ArrayList.ArrayList = ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list)
	fmt.Printf("list.TheSize: %v\n", list.TheSize)
}

func main5() {
	var list *ArrayList.ArrayList = ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list)
	list.Insert(1, "d")
	fmt.Println(list)
	for i := 0; i < 10; i++ {
		list.Insert(1, "x5")
		fmt.Println(list)
	}

	fmt.Printf("list.TheSize: %v\n", list.TheSize)
}

func main() {
	var list *ArrayList.ArrayList = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b1")
	list.Append("c1")
	list.Append("d1")
	list.Append("e1")
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		if item == "d1" {
			it.Remove()
		}
		fmt.Println(item)
	}
	fmt.Println(list)
}
