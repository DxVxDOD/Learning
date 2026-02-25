package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(namedReturn())
	fmt.Println(blancReturn())
	fmt.Println(nonesense())

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "*", "three"},
		{"2"},
	}
	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("Invalid length")
			continue
		}
		p1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := exp[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator")
			continue
		}
		p2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}

func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func namedReturn() (str string, n int) {
	str, n = "das", 10
	return str, n
}

func blancReturn() (str string, n int) {
	str, n = "What", 33
	return
}

func nonesense() (lol int) {
	return
}
