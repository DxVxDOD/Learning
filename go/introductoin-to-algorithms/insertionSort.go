package main

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
