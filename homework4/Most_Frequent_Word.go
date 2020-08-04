package main

import (
	"fmt"
	"strings"
)

func wordANDitsCount(input string) map[string]int {
	words := strings.Fields(input)
	counts := make(map[string]int)
	for _, word := range words {
		_, ok := counts[word]
		if ok {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func main() {
	fmt.Println("Program returns most frequent word in the text")
	input := "Go Python Java C++ PHP Go "

	var maxCount int
	var maxWord string
	for word, count := range wordANDitsCount(input) {
		maxCount = count
		maxWord = word
		break
	}
	for word, count := range wordANDitsCount(input) {
		if count > maxCount {
			maxCount = count
			maxWord = word
		}

	}
	fmt.Println(maxWord, "=>", maxCount)

}
