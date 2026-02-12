package main

import "fmt"

func main() {
	ex1()
	ex2()
}

func ex1() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")

	y := x[:2]
	z := x[2:]

	fmt.Println(cap(x), cap(y), cap(z))

	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
}

func ex2() {
	type Employee struct {
		firstName string
		lastName  string
		id        int64
	}

	var one Employee
	one.firstName = "D"
	one.lastName = "O"
	one.id = 134534
	fmt.Println(one)

	two := Employee{}
	two.id = 31232
	two.lastName = "R"
	two.firstName = "B"
	fmt.Println(two)

	three := Employee{
		"A",
		"O",
		33123,
	}
	fmt.Println(three)
}
