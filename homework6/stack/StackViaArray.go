package stack

import "fmt"

var stack [100]int
var n = 100
var top = -1

func Push(value int) bool {

	if top >= n-1 {
		fmt.Println("Stack overflow")
		return false
	}
	top++
	stack[top] = value
	return true

}
func Pop() (int, bool) {
	if top <= -1 {
		fmt.Println("Stack is empty")
		return 0, false
	}
	fmt.Println("Poped element is", stack[top])
	if top == 0 {
		top--
		return 0, true

	}
	top--

	return stack[top], true

}
