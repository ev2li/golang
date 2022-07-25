package main

import (
	"fmt"
	"sort"
)

func test_sort() {
	i := []int{2, 4, 1, 3}
	sort.Ints(i)
	fmt.Printf("i: %v\n", i)
}

type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func main() {
	n := []uint{1, 3, 2}
	sort.Sort(NewInts(n))
	fmt.Println(n)
}
