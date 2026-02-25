package main

import (
	"fmt"
	"math/rand"
)

func main() {
	varInIf()
	skipOuter()
	illegalGOTO()
	arr := ex1()
	ex2(arr)
	ex3()
}

func ex3() {
	var total int
	for i := range 10 {
		// Shadowed total
		// total := total + i
		total = total + i
	}
	fmt.Println(total)
}

func ex2(arr []int) {
	for _, n := range arr {
		switch {
		case n%2 == 0 && n%3 == 0:
			fmt.Println("Six!")
		case n%2 == 0:
			fmt.Println("Two!")
		case n%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}

func ex1() []int {
	var randomInts []int
	for range 100 {
		randomInts = append(randomInts, rand.Intn(100))
	}
	return randomInts
}

func illegalGOTO() {
	a := 10
	// goto skip
	b := 20
	// skip:
	c := 30
	fmt.Println(a, b, c)
	// if c > a {
	// 	goto inner
	// }
	if a < b {
		// inner:
		fmt.Println("a is smaller than b")
	}
}

func skipOuter() {
	strs := []string{"LOOOOOOOOOOL", "BIIIIG"}
outer:
	for _, v := range strs {
		for i, strV := range v {
			fmt.Println(i, strV, string(strV))
			if strV == 'O' {
				continue outer
			}
		}
	}
}

func varInIf() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("LOOOOOOOOOOL")
	} else if n < 5 {
		fmt.Println("SMAAAL")
	} else {
		fmt.Println("BIIIIG")
	}
}
