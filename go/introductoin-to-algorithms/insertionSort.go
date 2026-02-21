package main

import "fmt"

func main() {
	unSorted := []int{1, 3, 654, 765, 34, 34534, 5345, 345, 34534, 5346, 776876, 7634, 53}
	fmt.Println(unSorted)
	fmt.Println(insertionSort(unSorted))
}

func insertionSort(arr []int) []int {
	arrLen := len(arr)
	for i := 2; i < arrLen; i++ {
		el := arr[i]
		j := i - 1

		for j > 0 && arr[j] > el {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = el
	}
	return arr
}
