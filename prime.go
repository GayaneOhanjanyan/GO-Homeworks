package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Println("Process checks if it is prime number or not")
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
