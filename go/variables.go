package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	arr := []int{100, 100, 99}
	ht := make(map[int]int)
	hs := make(map[int]int)

	for _, num := range arr {
		if _, v := hs[num]; !v {
			hs[num] = num
			fmt.Println(v)
			fmt.Println("not in")
		} else {
			fmt.Println(v)
			fmt.Println("is in")
		}
	}
	fmt.Println(arr)
	slices.Sort(arr)
	values := slices.Collect(maps.Values(hs))
	fmt.Println("v", values)
	for i, num := range values {
		ht[num] = i + 1
	}
	fmt.Println("a", arr)
	for i, num := range arr {
		arr[i] = ht[num]
	}
	fmt.Println("a", arr)
}
