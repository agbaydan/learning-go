package greetings

import "fmt"

func Hello(name string) string {
	// := declares and initializes a variable in one line
	// it auto detects type based on the value assigned
	// to delcare separately: var message string
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}
