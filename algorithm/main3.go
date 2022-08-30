package main

import (
	"algorithm/Queue"
	"fmt"
	"io/ioutil"
)

func main1() {
	q := Queue.NewQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)

	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
}

func main() {
	path := "./"
	files := []string{}

	q := Queue.NewQueue()
	q.EnQueue(path)

	for {
		path := q.DeQueue()
		if path == nil {
			break
		}

		read, _ := ioutil.ReadDir(path.(string)) //读取文件夹
		for _, fi := range read {
			if fi.IsDir() {
				fullDir := path.(string) + fi.Name() //构造新的路径
				q.EnQueue(fullDir)
			} else {
				fullDir := path.(string) + fi.Name() //构造新的路径
				files = append(files, fullDir)
			}
		}
	}

	fmt.Printf("files: %v\n", files)
}
