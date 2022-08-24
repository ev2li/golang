package ArrayList

type StatckArrayIterator interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsFull() bool          //是否满了
	IsEmpty() bool         //是否为空
}

type StackIterator struct {
	array    *ArrayList
	Iterator Iterator
}

func NewArrayListStackIterator() *StackIterator {
	stack := new(StackIterator)
	stack.array = NewArrayList()
	stack.Iterator = stack.array.Iterator()
	return stack
}

func (stack *StackIterator) Clear() {
	stack.array.Clear()
	stack.array.TheSize = 0
}

func (stack *StackIterator) Size() int {
	return stack.array.TheSize
}

func (stack *StackIterator) Pop() interface{} {
	if !stack.IsEmpty() {
		last := stack.array.dataStrore[stack.array.TheSize-1]
		stack.array.Delete(stack.array.TheSize - 1)
		return last
	}
	return nil
}

func (stack *StackIterator) Push(data interface{}) {
	if !stack.IsFull() {
		stack.array.Append(data)
	}
}

func (stack *StackIterator) IsFull() bool {
	return stack.array.TheSize == 10
}

func (stack *StackIterator) IsEmpty() bool {
	return stack.array.TheSize == 0
}
