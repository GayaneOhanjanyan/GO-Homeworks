package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Println("Program checks if number is prime or not")
	fmt.Println("Input Number")
	n, err := fmt.Scan(&num)
	if n < 1 {
		fmt.Println("There is smt wrong with input, ", err)
		return
	}
	if !(num > 0) {
		fmt.Println("There is smt wrong with input, natural number is needed")
		return
	}
	if num == 1 {
		fmt.Println("It is not prime number")
		return
	}
	num1 := float64(num)
	root := math.Sqrt(num1)
	root1 := int(root)
	for i := 2; i <= root1; i++ {
		if num%i == 0 {
			fmt.Println(num, "is not prime number")
			return
		}
	}
	fmt.Println(num, "is prime number")
}
