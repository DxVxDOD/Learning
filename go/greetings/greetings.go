package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name was given")
	}
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, val := range names {
		message, err := Hello(val)
		if err != nil {
			return nil, err
		}
		messages[val] = message
	}

	return messages, nil
}

func randomGreeting() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
