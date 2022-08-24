package ArrayList

type StatckArray interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsFull() bool          //是否满了
	IsEmpty() bool         //是否为空
}

type Stack struct {
	array   *ArrayList
	capSize int //最大范围
}

func NewArrayListStack() *Stack {
	stack := new(Stack)
	stack.array = NewArrayList()
	stack.capSize = 10
	return stack
}

func (stack *Stack) Clear() {
	stack.array.Clear()
}

func (stack *Stack) Size() int {
	return stack.array.TheSize
}

func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty() {
		last := stack.array.dataStrore[stack.array.TheSize-1]
		stack.array.Delete(stack.array.TheSize - 1)
		return last
	}
	return nil
}

func (stack *Stack) Push(data interface{}) {
	if !stack.IsFull() {
		stack.array.Append(data)
	}
}

func (stack *Stack) IsFull() bool {
	return stack.array.TheSize == stack.capSize
}

func (stack *Stack) IsEmpty() bool {
	return stack.array.TheSize == 0
}
