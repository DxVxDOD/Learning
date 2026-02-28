package main

import (
	"errors"
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
	// cat()
	// fileLen()
	prefix := prefixer("Hello")
	fmt.Println(prefix("Puffy"))
}

func prefixer(pref string) func(str string) string {
	return func(str string) string {
		return fmt.Sprintf("%v, %v!", pref, str)
	}
}

func getFile(name string) (*os.File, func() error, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() error {
		return file.Close()
	}, nil
}

func fileLen() {
	if len(os.Args) < 2 {
		log.Fatal("Not eenough arguments")
	}
	file, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := closer(); err != nil {
			log.Fatal(err)
		}
	}()

	// var data []byte
	data := make([]byte, 2048)
	for {
		count, err := file.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		int, err := os.Stdout.Write(data[:count])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Stdout Write retuned val: ", int)
	}
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
		{"2", "/", "0"},
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
		// p3, err := strconv.Atoi(exp[3])
		// if err != nil {
		// 	fmt.Println(err)
		// 	continue
		// }
		// if p3 == 0 {
		// 	fmt.Println("Cannot devide by 0")
		// 	continue
		// }
		result, err = opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return result
}

func add(i, j int) (int, error) {
	return i + j, nil
}

func sub(i, j int) (int, error) {
	return i - j, nil
}

func mul(i, j int) (int, error) {
	if i == 0 || j == 0 {
		return 0, errors.New("mult by 0")
	}
	return i * j, nil
}

func div(i, j int) (int, error) {
	if j == 0 || i == 0 {
		return 0, errors.New("div by 0")
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
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
