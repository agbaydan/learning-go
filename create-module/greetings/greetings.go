package greetings

import (
	"fmt"
	
	"errors"
)

func Hello(name string) (string, error) {
	// := declares and initializes a variable in one line
	// it auto detects type based on the value assigned
	// to delcare separately: var message string

	if (name == "") {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	
	// nil means no error aka empty value
	return message, nil
}
