package greetings

import (
	"fmt"
	
	"errors"

	"math/rand"

	"time"
)

func Hello(name string) (string, error) {
	// := declares and initializes a variable in one line
	// it auto detects type based on the value assigned
	// to delcare separately: var message string

	if (name == "") {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	
	// nil means no error aka empty value
	return message, nil
}

// https://go.dev/doc/effective_go#init
// Go calls init functions at program startup after global vars have been initialized
// so kinda like a ctor for this file
func init() {
	rand.Seed(time.Now().UnixNano())
}

// starting with lowercase means this method is essentially internal.
// it is not exported and not accessible by other packages
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"You and I could rule this city, %v. Or, we could fight to the death!",
	}

	return formats[rand.Intn(len(formats))]
}
