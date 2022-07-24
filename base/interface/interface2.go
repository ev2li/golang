package main

import "fmt"

// 一个类型实现多个接口
type Player interface {
	PlayMusic()
}

type PlayVideo interface {
	PlayVideo()
}

type Mobile01 struct {
}

func (mobile Mobile01) PlayMusic() {
	fmt.Println("play musix")
}

func (mobile Mobile01) PlayVideo() {
	fmt.Println("play video")
}

type Pet01 interface {
	eat01()
}

type Dog01 struct {
}

type Cat01 struct {
}

func (dog Dog01) eat01() {
	fmt.Println("dog eat...")
}

func (cat Cat01) eat01() {
	fmt.Println("cat eat...")
}

func main() {
	m := Mobile01{}
	m.PlayVideo()
	m.PlayMusic()

	/* 	dog01 := Dog01{}
	   	cat01 := Cat01{}
	   	dog01.eat01()
	   	cat01.eat01() */

	var pet01 Pet01
	pet01 = Dog01{}
	pet01.eat01()
	pet01 = Cat01{}
	pet01.eat01()
}
