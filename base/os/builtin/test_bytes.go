package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func testTranse() {
	//强制类型转换
	var i int = 100
	var b byte = 10
	b = byte(i)
	i = int(b)

	var s string = "hello world"
	b1 := []byte{1, 2, 3}
	s = string(b1)
	b1 = []byte(s)
}

func testContains() {
	s := "ev2li.com"
	b := []byte(s)
	b1 := []byte("ev2li")
	b2 := []byte("Ev2li")
	fmt.Printf("strings.Contains(\"hello world\", \"hello\"): %v\n", strings.Contains("hello world", "hello"))
	b3 := bytes.Contains(b, b1)
	fmt.Printf("b3: %v\n", b3)
	b3 = bytes.Contains(b, b2)
	fmt.Printf("b3: %v\n", b3)
}

func testCount() {
	s := []byte("helloooooooooo")
	sep1 := []byte("h")
	sep2 := []byte("l")
	sep3 := []byte("o")
	fmt.Println(bytes.Count(s, sep1))
	fmt.Println(bytes.Count(s, sep2))
	fmt.Println(bytes.Count(s, sep3))
}

func testReplace() {
	s := []byte("hello world")
	old := []byte("o")
	new := []byte("ee")
	fmt.Println(string(bytes.Replace(s, old, new, 0)))  //hello world
	fmt.Println(string(bytes.Replace(s, old, new, 1)))  //hellee world
	fmt.Println(string(bytes.Replace(s, old, new, 2)))  //hellee weerld
	fmt.Println(string(bytes.Replace(s, old, new, -1))) //hellee,weerld
}

func testRunes() {
	s := []byte("你好世界")
	r := bytes.Runes(s)
	fmt.Println("转换前字符串长度", len(s))
	fmt.Println("转换后字符串长度", len(r))
}

func testJoin() {
	s2 := [][]byte{[]byte("你好"), []byte("世界")}
	sep4 := []byte(",")
	fmt.Println(string(bytes.Join(s2, sep4))) //你好,世界
	sep5 := []byte("#")
	fmt.Println(string(bytes.Join(s2, sep5))) //你好#世界

}

func testReader() {
	data := "12345678"
	r := bytes.NewReader([]byte(data))
	fmt.Printf("r.Len(): %v\n", r.Len())

}

func testBytesBuffer() {
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)
	var b1 = bytes.NewBufferString("hello")
	fmt.Printf("b1: %v\n", b1)
	var b2 = bytes.NewBuffer([]byte("hello"))
	fmt.Printf("b2: %v\n", b2)
}
func testBytesBuffer2() {
	var b bytes.Buffer
	n, _ := b.WriteString("hello")
	fmt.Println(n)
	fmt.Printf("b: %v\n", string(b.Bytes()))

}
func testBytesBuffer3() {
	var b = bytes.NewBufferString("hello world")
	b1 := make([]byte, 2)
	for {
		n, err := b.Read(b1)
		if err == io.EOF {
			break
		}
		fmt.Println(n)
		fmt.Printf("b1: %v\n", string(b1[0:n]))
	}
}

func main() {
	// testCount()
	// testReplace()
	// testRunes()
	// testJoin()
	// testReader()
	// testBytesBuffer()
	// testBytesBuffer2()
	testBytesBuffer3()
}
