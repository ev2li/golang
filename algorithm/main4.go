package main

import (
	"algorithm/Queue"
	"fmt"
)

func main() {
	var cq Queue.CircleQueue
	Queue.InitQueue(&cq)
	Queue.EnQueue(&cq, 1)
	Queue.EnQueue(&cq, 2)
	Queue.EnQueue(&cq, 3)
	Queue.EnQueue(&cq, 4)
	Queue.EnQueue(&cq, 5)

	fmt.Println(Queue.DeQueue(&cq))
	fmt.Println(Queue.DeQueue(&cq))
	fmt.Println(Queue.DeQueue(&cq))
	fmt.Println(Queue.DeQueue(&cq))

}
