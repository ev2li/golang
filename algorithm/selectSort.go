package main

import (
	"fmt"
	"strings"
)

func SelctSortMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0] //一个元素的数组，直接返回
	} else {
		max := arr[0] //假定第一个最大
		for i := 1; i < length; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}

func main1() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	i := SelctSortMax(arr)
	fmt.Printf("i: %v\n", i)
	newArr := SelctSort(arr)
	fmt.Println(newArr)
}

func SelctSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		for i := 0; i < length-1; i++ {
			min := i
			for j := i + 1; j < length; j++ { //每次选出一个最小值
				if arr[min] < arr[j] {
					min = j
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] //数据交换
			}
		}
		return arr
	}
}

func StrSort(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		for i := 0; i < length-1; i++ {
			min := i
			for j := i + 1; j < length; j++ { //每次选出一个最小值
				/*if arr[min] < arr[j] {
					min = j
				}*/
				if strings.Compare(arr[min], arr[j]) < 0 {
					min = j
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] //数据交换
			}
		}
		return arr
	}
}

func StrSortMax(arr []string) string {
	length := len(arr)
	if length <= 1 {
		return arr[0] //一个元素的数组，直接返回
	} else {
		max := arr[0] //假定第一个最大
		for i := 1; i < length; i++ {
			/*if arr[i] > max {
				max = arr[i]
			}*/
			if strings.Compare(arr[i], max) > 0 {
				max = arr[i]
			}

		}
		return max
	}
}

func main2() {
	arr := []string{"c", "a", "b", "x", "z", "m", "n", "d", "f"}
	fmt.Println(StrSortMax(arr))
	fmt.Println(StrSort(arr))
}

func BubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		for i := 0; i < length-1; i++ {
			for j := 0; j < length-1-i; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
		return arr
	}
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	fmt.Println(BubbleSort(arr))
}
