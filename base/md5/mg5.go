package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	//方法一
	data := []byte("test")
	s := fmt.Sprintf("%x", md5.Sum(data))
	fmt.Println(s)

	//方法二:
	b := []byte("test")
	h := md5.New()
	h.Write(b)
	fmt.Printf("hex.EncodeToString(h.Sum(nil)): %v\n", hex.EncodeToString(h.Sum(nil)))
	//加盐
	salt := "ev2li"
	b2 := []byte(salt)
	h.Write(b2)
	fmt.Printf("hex.EncodeToString(h.Sum(nil)): %v\n", hex.EncodeToString(h.Sum(nil)))
}
