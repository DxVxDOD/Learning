package main

import "fmt"

func main() {
	fmt.Println("Hello Mom and Dad")
	test := [...]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	containsDuplicate(test[:])
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, isIn := m[nums[i]]
		println("IS IN", isIn)
		println()
		if !isIn {
			m[nums[i]] = i
		} else {
			return true
		}
	}
	for k, v := range m {
		println(k, v)
	}
	return false
}
