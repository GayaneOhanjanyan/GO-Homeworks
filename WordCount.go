package main

import "fmt"

func wordCount(input string) int {

	count := 0
	s := []rune(input)

	for i := 1; i < len(s)-1; i++ {

		if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			if len(s) == 1 {
				return 0
			}
			count++
		}
	}
	count++
	return count
}
func main() {
	input := "The first programmer in the world was a woman. Her name was Ada Lovelace."
	fmt.Println("Program returns word count in the given string")
	c := wordCount(input)
	fmt.Println(c)
}
