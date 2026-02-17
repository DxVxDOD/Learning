package main

import "fmt"

// type DynamicArray struct {
// 	arr []int
// }

// func newDynamicArray(cap int) *DynamicArray {
// 	da := &DynamicArray{}
// 	da.arr = make([]int, 0, cap-2)
// 	val := da.arr[len(da.arr)-2]
// 	newArr := make([]int, 0, cap-2)
// 	return da
// }

func kthDistinct(k int) string {
	arr := []string{"d", "b", "c", "b", "c", "a"}
	var chars [26]int
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		chars[[]byte(arr[i])[0]-byte('a')]++
	}
	for i := 0; i < arrLen; i++ {
		currChar := chars[[]byte(arr[i])[0]-byte('a')]
		if currChar == 1 {
			k--
		}
		if k == 0 && currChar == 1 {
			return arr[i]
		}

		fmt.Printf("K: %d ;", k)
	}
	fmt.Printf("The charArr: %v", chars)
	return "denada"
}

func main() {
	res := kthDistinct(2)
	fmt.Println(res)
}
