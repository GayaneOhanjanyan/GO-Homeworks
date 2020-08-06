package main

import (
	"./stackinterface"
)

func main() {
	var str stackinterface.ArrayStack
	str.Push(10)
	str.Push(12)
	str.PrintStack()
	str.Pop()

}
