package main

import "fmt"

type New interface {
	Init(value ...interface{}) any
}

type Add struct {
	Num int
}

func (add *Add) Init(num ...interface{}) any {
	add.Num = num[0].(int)
	return *add
}

type Sub struct {
	Num int
}

func (sub *Sub) Init(num ...interface{}) any {
	sub.Num = num[0].(int)
	return *sub
}

func main() {
	var (
		b = New(&Sub{}).Init(20).(Sub)
		a = New(&Add{}).Init(10).(Add)
	)
	fmt.Println(a, b)
}
