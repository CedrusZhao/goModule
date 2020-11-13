package myFirst

import "fmt"

func MyName() {
	fmt.Print("zhaos")
}

type people1 struct {
	name string
	age  int
}

type People2 struct {
	name string
	age  int
}

func learn1() {
	fmt.Println("learn1")
}

func Learn2() {
	fmt.Println("learn2")
}

type read1 interface {
	read1()
}

type Read2 interface {
	Read2()
}
