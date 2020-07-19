package main

import "fmt"

func main() {
	array := []int{10, 25, 35, 0, -5}
	max, index := max(array)
	fmt.Println(max, index)
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
