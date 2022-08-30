package Link

import "errors"

type Node struct {
	Data  interface{}
	PNext *Node
}

type LinkStack interface {
	IsEmpty() bool //是否为空
	Push(data any)
	Pop() any
	Length() int
}

func NewStack() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool {
	return n.PNext == nil
}

//头插法
func (n *Node) Push(data any) {
	newNode := &Node{Data: data, PNext: nil}
	newNode.PNext = n.PNext
	n.PNext = newNode
}

//头删法
func (n *Node) Pop() (any, error) {
	if n.IsEmpty() {
		return nil, errors.New("Bug")
	}

	value := n.PNext.Data //要弹出的数据
	n.PNext = n.PNext.PNext
	return value, nil
}

func (n *Node) Length() int {
	pnext := n
	length := 0
	for pnext.PNext != nil {
		pnext = pnext.PNext
		length++
	}
	return length
}
