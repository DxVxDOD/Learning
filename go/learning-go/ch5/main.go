package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type Person struct {
	fName string
	lName string
	age   int
}

func main() {
	fmt.Println(namedReturn())
	fmt.Println(blancReturn())
	fmt.Println(nonesense())
	fmt.Println(calculator())
	fmt.Println(anonymus())
	clousreInSort()
	cat()
}

func cat() {
	if len(os.Args) < 2 {
		log.Fatal("No file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func clousreInSort() {
	people := []Person{
		{"Puffy", "Puff", 27},
		{"Puf", "Pu", 26},
		{"Puffu", "Puffulet", 25},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].fName < people[j].fName
	})
	fmt.Println("Sorted by first name: ", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println("Sorted by last age: ", people)
}

func anonymus() int {
	f := func(a int) int {
		return a * a
	}
	return f(4)
}

func calculator() int {
	var result int
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
		result = opFunc(p1, p2)
	}
	return result
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
