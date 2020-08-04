package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 int
	fmt.Println("Program outputs the prime factors of the number")
	fmt.Println("Input Number")
	n, err := fmt.Scan(&num1)
	if n < 1 {
		fmt.Println("There is smt wrong with input, ", err)
		return
	}
	if !(num1 > 0) {
		fmt.Println("There is smt wrong with input, natural number is needed")
		return
	}
	if num1 == 1 {
		fmt.Println("It does not have any prime factors")
		return
	}
	fmt.Println("The prime factor(s) of", num1, "is(are):")
	fmt.Print(num1, "=")
	var count int
	num := num1
	for num%2 == 0 {
		count++
		num = num / 2
	}
	if count > 1 {
		fmt.Print("2^", count)
	} else if count == 1 {
		fmt.Print("2")
	}
	num2 := float64(num)
	root := math.Sqrt(num2)
	root1 := int(root)

	for i := 3; i <= root1; i = i + 2 {
		count := 0

		for num%i == 0 {

			count++
			num = num / i

		}
		if count > 1 {
			fmt.Print("*", i, "^", count)
		} else if count == 1 {
			fmt.Print("*", i)

		}

	}

	if num > 2 && num1%2 == 0 {
		fmt.Print("*", num)
	}
	if num > 2 && !(num1%2 == 0) {
		fmt.Print(num)
	}

}
