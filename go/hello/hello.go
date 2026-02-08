package main

import "fmt"

func main() {
	calcExp := func(x, z int64) int64 {
		var y int64 = 1

		for z != 0 {
			r := z % 2
			z = z / 2
			if r == 1 {
				y = x * y
			}
			x = x * x
		}
		return y
	}

	res := calcExp(2, 25)

	fmt.Println(res)
}
