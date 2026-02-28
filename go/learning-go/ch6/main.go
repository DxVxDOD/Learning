package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		firstName,
		lastName,
		age,
	}
}

func MakePersonPointer(firstNamePointer, lastNamePointer string, agePointer int) *Person {
	return &Person{
		firstNamePointer,
		lastNamePointer,
		agePointer,
	}
}

func UpdateSlice(strSlice []string, str string) {
	strSlice[len(strSlice)-1] = str
	fmt.Println(strSlice)
}

func GrowSlice(strSlice []string, str string) {
	strSlice = append(strSlice, str)
	fmt.Println(strSlice)
}

func build10K() {
	sl := make([]Person, 0, 10000000)
	// for range 10000000 {
	// 	sl = append(sl, MakePerson("Puffy", "Puff", 28))
	// }
	fmt.Println(cap(sl))
}

func main() {
	// MakePerson("Puff", "Puffy", 28)
	// MakePersonPointer("Puff", "Puffy", 28)
	// strSlice := []string{"Puff", "Puffy"}
	// fmt.Println("Start", strSlice)
	// UpdateSlice(strSlice, "Pufu")
	// fmt.Println("After UpdateSlice", strSlice)
	// GrowSlice(strSlice, "Pufulet")
	// fmt.Println("After GrowSlice", strSlice)
	build10K()
}
