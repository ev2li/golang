package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool //是否有下一个
	Next() (interface{}, error)
	Remove()       //删除
	GetIndex() int //得到索引
}

type Iterable interface {
	Iterator() Iterator //构造初始化接口
}

// 构造指针访问数组
type ArrayListIterator struct {
	list         *ArrayList //数组指针
	currentIndex int        //当前索引

}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)
	it.currentIndex = 0
	it.list = list
	return it
}

func (it *ArrayListIterator) HasNext() bool {
	return it.currentIndex < it.list.TheSize //是否有下一个
}

func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("has not next")
	}

	value, err := it.list.Get(it.currentIndex) //抓取当前数据
	it.currentIndex++
	return value, err
}

func (it *ArrayListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex) //删除一个元素
}

func (it *ArrayListIterator) GetIndex() int {
	return it.currentIndex
}
