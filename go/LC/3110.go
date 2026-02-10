package main

import (
	"fmt"
	"math"
)

func main() {
	s := "tzt"
	sLen := len(s)
	charCount := make([]int, 26)
	for i := range sLen {
		charCount[s[i]-byte('a')] += 1
	}
	fmt.Println(charCount)
	minV, maxV := math.MaxInt, 0
	for i := range sLen {
		sEl := charCount[i]
		fmt.Println(sEl % 2)
		if sEl > 1 && (sEl%2) == 0 {
			minV = min(sEl, minV)
			fmt.Println(minV)
		} else {
			maxV = max(sEl, maxV)
		}
	}

	fmt.Println(maxV - minV)
}
