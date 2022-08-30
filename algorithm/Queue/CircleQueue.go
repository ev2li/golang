package Queue

import "errors"

const QueueSize = 100

type CircleQueue struct {
	dataStore [QueueSize]interface{} //队列的数据存储
	front     int                    //头部位置
	rear      int                    //尾部位置
}

//初始化，头部和尾部重合为空
func InitQueue(q *CircleQueue) {
	q.front = 0
	q.rear = 0
}

func QueueLength(q *CircleQueue) int {
	return ((q.rear - q.front + QueueSize) % QueueSize)
}

func EnQueue(q *CircleQueue, data interface{}) (err error) {
	if (q.rear+1)%QueueSize == q.front%QueueSize {
		return errors.New("队列已经满了")
	}
	q.dataStore[q.rear] = data //入队
	q.rear = (q.rear + 1) % QueueSize
	return nil
}

func DeQueue(q *CircleQueue) (data interface{}, err error) {
	if q.front == q.rear {
		return nil, errors.New("Queue is empty")
	}

	res := q.dataStore[q.front] //取出第一个元素
	q.dataStore[q.front] = 0    //清空数据
	q.front = (q.front + 1) % QueueSize

	return res, nil
}
