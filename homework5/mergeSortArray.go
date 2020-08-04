package main

import "fmt"

func mergeArray(arr1 []int, arr2 []int) (arr3 [50]int) {
	i := 0
	j := 0
	k := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr3[k] = arr1[i]
			k++
			i++
		} else {
			arr3[k] = arr2[j]
			k++
			j++
		}
	}
	for i < len(arr1) {
		arr3[k] = arr1[i]
		k++
		i++
	}
	for j < len(arr2) {
		arr3[k] = arr2[j]
		k++
		j++
	}
	return
}
func sort(arr1 []int) []int {
	for j := 0; j < len(arr1); j++ {
		for i := 0; i < len(arr1); i++ {
			if arr1[i] > arr1[j] {
				k := arr1[j]
				arr1[j] = arr1[i]
				arr1[i] = k
			}
		}
	}
	return arr1
}
func main() {
	fmt.Println("Program merges given 2 arrays into one big sorted array and returns it")
	arr1 := []int{3, -5, 9, 12, 16, 17, -100}
	sArray1 := sort(arr1)
	fmt.Println("Sorted array1=>", sArray1)
	arr2 := []int{1, -3, 5, 10}
	sArray2 := sort(arr2)
	fmt.Println("Sorted array2=>", sArray2)
	arr3 := mergeArray(sArray1, sArray2)
	fmt.Print("Merged array is=> ")
	for i := 0; i < len(sArray1)+len(sArray2); i++ {
		fmt.Print(arr3[i], " ")
	}
}
