package main

import (
	"fmt"
	"io/ioutil"
)

func GetAllNew(path string, level int) {
	fmt.Println("level:", level)
	levelStr := ""
	if level == 1 {
		levelStr = "+"
	} else {
		for ; level > 1; level-- {
			levelStr += "|--"
		}
		levelStr += "+"
	}
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return
	}
	for _, fi := range read { //循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fullDir := path + fi.Name() //构造新的路径
			fmt.Println(levelStr + fullDir)
			GetAllNew(fullDir, level+1) //文件夹递归处理
		} else {
			fullDir := path + fi.Name() //构造新的路径
			fmt.Println(levelStr + fullDir)
		}
	}
}

func main() {
	path := "./" // 路径
	//files := []string{} // 数组字符串

	GetAllNew(path, 1)
}
