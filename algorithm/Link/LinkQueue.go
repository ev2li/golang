package Link

type QueueLink struct {
	Rear  *Node
	Front *Node
}

type LinkQueue interface {
	Length() int
	EnQueue(value any)
	DeQueue() any
}

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (qlk *QueueLink) Length() int {
	pnext := qlk.Front
	length := 0
	for pnext.PNext != nil {
		pnext = pnext.PNext
		length++
	}
	return length
}

func (qlk *QueueLink) EnQueue(value any) {
	newnode := &Node{value, nil}
	if qlk.Front == nil {
		qlk.Front = newnode
		qlk.Rear = newnode
	} else {
		qlk.Rear.PNext = newnode
		qlk.Rear = qlk.Rear.PNext
	}
}

func (qlk *QueueLink) DeQueue() any {
	if qlk.Front == nil {
		return nil
	}

	newnode := qlk.Front       //记录头部位置
	if qlk.Front == qlk.Rear { //只有一个
		qlk.Front = nil
		qlk.Rear = nil
	} else {
		qlk.Front = qlk.Front.PNext //删除一个
	}
	return newnode.Data
}
