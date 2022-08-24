package StackArray

type StatckArray interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsFull() bool          //是否满了
	IsEmpty() bool         //是否为空
}

type Stack struct {
	dataStore   []interface{}
	capSize     int //最大范围
	currentSize int //当前实际使用大小
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.dataStore = make([]interface{}, 0, 10)
	stack.capSize = 10
	stack.currentSize = 0
	return stack
}

func (stack *Stack) Clear() {
	stack.dataStore = make([]interface{}, 0, 10)
	stack.capSize = 10
	stack.currentSize = 0
}

func (stack *Stack) Size() int {
	return stack.currentSize
}

func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty() {
		last := stack.dataStore[stack.currentSize-1]            //最后一个数据
		stack.dataStore = stack.dataStore[:stack.currentSize-1] //删除最后一个
		stack.currentSize--                                     //删除
		return last
	}
	return nil
}

func (stack *Stack) Push(data interface{}) {
	if !stack.IsFull() {
		stack.dataStore = append(stack.dataStore, data) //压入
		stack.currentSize++
	}
}

func (stack *Stack) IsFull() bool {
	return stack.currentSize == stack.capSize
}

func (stack *Stack) IsEmpty() bool {
	return stack.currentSize == 0
}
