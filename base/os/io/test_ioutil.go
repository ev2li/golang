package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func testReadAll() {
	f, _ := os.Open("1.txt")
	defer f.Close()
	// r := strings.NewReader("Hello world")
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("string(b): %v\n", string(b))
}

func testReadDir() {
	//读取文件夹
	fi, _ := ioutil.ReadDir("./")
	for _, v := range fi {
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func testReadFile() {
	b, _ := ioutil.ReadFile("1.txt")
	fmt.Printf("b: %v\n", string(b))

}

func testWriteFile() {
	err := ioutil.WriteFile("test.txt", []byte("hello world"), os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func testCreateTmpFile() {
	f, err := ioutil.TempFile("./", "example")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer os.Remove(f.Name()) //clean up

	if _, err := f.Write([]byte("hello world")); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	if err := f.Close(); err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func main() {
	// testReadAll()
	// testReadDir()
	// testReadFile()
	// testWriteFile()
	testCreateTmpFile()
}
