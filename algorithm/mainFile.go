package main

import (
	"algorithm/StackArray"
	"errors"
	"fmt"
	"io/ioutil"
)

func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, fi := range read { //循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fullDir := path + fi.Name()       //构造新的路径
			files = append(files, fullDir)    //追加路径
			files, _ = GetAll(fullDir, files) //文件夹递归处理
		} else {
			fullDir := path + fi.Name()    //构造新的路径
			files = append(files, fullDir) //追加路径
		}
	}
	return files, nil
}

func main1() {
	path := "./"        // 路径
	files := []string{} // 数组字符串

	files, _ = GetAll(path, files)
	for i := 0; i < len(files); i++ { //打印路径
		fmt.Println(files[i])
	}
}

//使用栈遍历文件夹
func main() {
	path := "./"
	files := []string{}
	myStack := StackArray.NewStack()
	myStack.Push(path)
	for !myStack.IsEmpty() {
		path := myStack.Pop().(string)
		files = append(files, path)
		read, _ := ioutil.ReadDir(path) //读取文件夹下面所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fullDir := path + fi.Name() //构造新的路径
				// files = append(files, fullDir)
				myStack.Push(fullDir)
			} else {
				fullDir := path + fi.Name() //构造新的路径
				files = append(files, fullDir)
			}
		}
	}
	for i := 0; i < len(files); i++ { //打印路径
		fmt.Println(files[i])
	}
}
