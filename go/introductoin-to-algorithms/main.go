package main

import "fmt"

func main() {
	unSorted := []int{1, 3, 654, 765, 34, 34534, 5345, 345, 34534, 5346, 776876, 7634, 53}
	fmt.Println(unSorted)
	fmt.Println(insertionSort(unSorted))
}
