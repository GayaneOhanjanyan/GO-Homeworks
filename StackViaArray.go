package main

import "fmt"

var stack [100]int
var n = 100
var top = -1

func Push(value int) bool {

	if top >= n-1 {
		fmt.Println("Stack overflow")
		return false
	} else {
		top++
		stack[top] = value
		return true
	}

}
func Pop() (int, bool) {
	if top <= -1 {
		fmt.Println("Stack is empty")
		return 0, false
	} else {
		fmt.Println("Poped element is", stack[top])
		if top == 0 {
			top--
			return 0, true

		} else {
			top--
		}

		return stack[top], true

	}
}
func PrintSTACK() {
	if top >= 0 {
		fmt.Println("Stack elements are:")
		for i := top; i >= 0; i-- {
			fmt.Println(stack[i])
		}
	} else {
		fmt.Println("Stack is empty")
	}
}
func main() {
	var choice int
	var value int
	fmt.Println("1)Push in stack")
	fmt.Println("2)Pop from stack")
	fmt.Println("3)Print stack")
	fmt.Println("4)Exit")
	for {
		fmt.Println("Enter choice:")
		n, err := fmt.Scan(&choice)
		if n < 1 {
			fmt.Println("There is smt wrong with input, ", err)
			return
		}
		switch choice {
		case 1:
			{
				fmt.Println("Enter value to be pushed")
				n, err := fmt.Scan(&value)
				if n < 1 {
					fmt.Println("There is smt wrong with input, ", err)
					return
				}
				Push(value)
				break
			}
		case 2:
			{
				Pop()
				break
			}
		case 3:
			{
				PrintSTACK()
				break
			}
		default:
			{
				fmt.Println("Invalid Choice")
			}
		case 4:
			{
				fmt.Println("Exit")
				return
			}

		}
	}
}
