package main

import "fmt"

func main() {
	strs := []string{"ab", "a"}

	res := ""
	strsLen := len(strs)
	first := strs[0]
	firstLen := len(first)
	for i := 0; i < firstLen; i++ {
		fEl := first[i]
		for j := 0; j < strsLen; j++ {
			fmt.Println(i, j, strsLen)
			strsEl := strs[j][i]
			if strsEl != fEl {
				fmt.Println(res)
			}
		}
		res += string(fEl)
	}
	fmt.Println(res)

	fmt.Println(res)
}
