package ArrayList

import (
	"errors"
	"fmt"
)

// 接口
type List interface {
	Size() int                                  //数组大小
	Get(index int) (interface{}, error)         //抓取第几个元素
	Set(index int, newVal interface{}) error    //修改数据
	Insert(index int, newVal interface{}) error //插入数据
	Append(newVal interface{})                  //追加
	Clear()                                     //清空
	Delete(index int) error                     //删除
	String() string                             //返回字符串
}

type ArrayList struct {
	dataStrore []interface{} //数组的存储
	TheSize    int           //数组的大小
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)                       //初始化结构体
	list.dataStrore = make([]interface{}, 0, 10) //开辟空间10个
	list.TheSize = 0
	return list
}

// 数组的大小
func (list *ArrayList) Size() int {
	return list.TheSize
}

// 抓取数据
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index > list.TheSize {
		return nil, errors.New("index out of range")
	}
	return list.dataStrore[index], nil
}

// 插入数据
func (list *ArrayList) Append(newVal interface{}) {
	list.dataStrore = append(list.dataStrore, newVal) //添加数据
	list.TheSize++
}

// 输出数组
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStrore)
}

// 修改数据
func (list *ArrayList) Set(index int, newVal interface{}) error {
	if index < 0 || index > list.TheSize {
		return errors.New("index out of range")
	}
	list.dataStrore[index] = newVal //设置新值
	return nil
}

// 删除数据
func (list *ArrayList) Delete(index int) error {
	if index < 0 || index > list.TheSize {
		return errors.New("index out of range")
	}
	list.dataStrore = append(list.dataStrore[:index], list.dataStrore[index+1:]...)
	list.TheSize--

	return nil
}

//插入数据
func (list *ArrayList) Insert(index int, newVal interface{}) error {
	if index < 0 || index > list.TheSize {
		return errors.New("index out of range")
	}
	list.checkIsFull() //监测内存，如果满了，扩容
	list.dataStrore = list.dataStrore[:list.TheSize+1]
	for i := list.TheSize; i > index; i-- {
		list.dataStrore[i] = list.dataStrore[i-1] //往后移动
	}
	list.dataStrore[index] = newVal //插入数据
	list.TheSize++
	return nil
}

// 扩容
func (list *ArrayList) checkIsFull() {
	if list.TheSize == cap(list.dataStrore) {
		//扩容
		newDataStore := make([]interface{}, 2*list.TheSize)
		copy(newDataStore, list.dataStrore)
		/*for i := 0; i < list.TheSize; i++ {
			newDataStore[i] = list.dataStrore[i]
		}*/
		list.dataStrore = newDataStore
	}
}

// 清空数据
func (list *ArrayList) Clear() {
	list.dataStrore = make([]interface{}, 0, 10)
	list.TheSize = 0
}
