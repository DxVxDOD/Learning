package main

type DynamicArray struct {
	arr []int
}

func newDynamicArray(cap int) *DynamicArray {
	da := &DynamicArray{}
	da.arr = make([]int, 0, cap-2)
	val := da.arr[len(da.arr)-2]
	newArr := make([]int, 0, cap-2)
	return da
}

func main() {
}
