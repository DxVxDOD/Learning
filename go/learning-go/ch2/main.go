package main

import (
	"fmt"
	"math"
)

func main() {
	ex1()
	ex2()
	ex3()
}

func ex1() {
	i := 20
	f := float32(i)
	fmt.Println(i)
	fmt.Println(f)
}

func ex2() {
	const value = 10
	var i int
	var f float32
	i = value
	f = value

	fmt.Println(i)
	fmt.Println(f)
}

func ex3() {
	var b byte
	var smalI int32
	var bigI uint64

	b = math.MaxUint8
	fmt.Println(b + 1)

	smalI = math.MaxInt32
	fmt.Println(smalI + 1)

	bigI = math.MaxUint64
	fmt.Println(bigI + 1)
}
