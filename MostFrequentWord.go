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
func max(array []int) (int, int) {
	maxel := array[0]
	index := 0
	for i := 1; i < len(array); i++ {
		if maxel < array[i] {
			maxel = array[i]
			index = i
		}
	}
	return maxel, index
}
func main() {
	fmt.Println("Program returns most frequent word in the text")
	input := "Go Python C++ PHP Java Go"

	var c []int
	var w []string
	for word, count := range wordANDitsCount(input) {

		c = append(c, count)
		w = append(w, word)

	}

	co, index := max(c)

	fmt.Println(w[index], "=>", co)

}
