package Queue

type QueueI interface {
	Size() int                //大小
	Front() interface{}       //第一个元素
	End() interface{}         //最后一个元素
	IsEmpty() bool            //是否为空
	EnQueue(data interface{}) //入队
	DeQueue() interface{}     //出队
	Clear()                   //清空
}

type Queue struct {
	dataStore []interface{} //队列的数据存储
	theSize   int           //队列的大小
}

func NewQueue() *Queue {
	queue := new(Queue) //初始化，开辟结构体
	queue.dataStore = make([]interface{}, 0, 1000)
	queue.theSize = 0
	return queue
}

//大小
func (queue *Queue) Size() int {
	return queue.theSize
}

//第一个元素
func (queue *Queue) Front() interface{} {
	if queue.Size() == 0 {
		return nil
	}

	return queue.dataStore[0]
}

//最后一个元素
func (queue *Queue) End() interface{} {
	if queue.Size() == 0 {
		return nil
	}
	return queue.dataStore[queue.Size()-1]
}

//是否为空
func (queue *Queue) IsEmpty() bool {
	return queue.theSize == 0

}

//入队
func (queue *Queue) EnQueue(data interface{}) {
	queue.dataStore = append(queue.dataStore, data) //入队
	queue.theSize++
}

//出队
func (queue *Queue) DeQueue() interface{} {
	if queue.Size() == 0 {
		return nil
	}

	data := queue.dataStore[0]
	if queue.Size() > 1 {
		queue.dataStore = queue.dataStore[1:queue.Size()]
	} else {
		queue.dataStore = make([]interface{}, 0, 1000)
	}
	queue.theSize--
	return data
}

//清空
func (queue *Queue) Clear() {
	queue.dataStore = make([]interface{}, 0, 1000)
	queue.theSize = 0
}
