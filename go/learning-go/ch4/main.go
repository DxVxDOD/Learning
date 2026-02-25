package main

import (
	"fmt"
	"math/rand"
)

func main() {
	varInIf()
	skipOuter()
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
