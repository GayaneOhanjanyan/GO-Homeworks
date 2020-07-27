package main

import (
	"fmt"

	"./stack"
)

func main() {
	var choice int
	var value int
	fmt.Println("1)Push in stack")
	fmt.Println("2)Pop from stack")
	fmt.Println("3)Exit")
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
				stack.Push(value)
				break
			}
		case 2:
			{
				stack.Pop()
				break
			}

		default:
			{
				fmt.Println("Invalid Choice")
			}
		case 3:
			{
				fmt.Println("Exit")
				return
			}

		}
	}
}
