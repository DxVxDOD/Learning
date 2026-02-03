package main

import "fmt"

func main() {
	x := 200
	var b byte = 255

	fmt.Println("x: ", x, "b: ", b)

	sum1 := x + int(b)
	sum2 := byte(x) + b

	fmt.Println("byte(x): ", byte(x), "int(b): ", int(b))

	fmt.Println("sum1: ", sum1, "sum2: ", sum2)
}
