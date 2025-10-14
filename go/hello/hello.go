package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	messageOne, err := greetings.Hello("Puffy")
	if err != nil {
		log.Fatal(err)
	}

	messageTwo, err := greetings.Hello("Ari")
	if err != nil {
		log.Fatal(err)
	}

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messageOne)
	fmt.Println(messageTwo)

	for _, val := range messages {
		fmt.Println(val)
	}
}
