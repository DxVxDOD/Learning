package main

import "fmt"

func main() {
	det := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}
	res := 0
	for _, el := range det {
		age := el[11:13]
		fmt.Println("Age", age)
	}

	fmt.Println(res)
}
