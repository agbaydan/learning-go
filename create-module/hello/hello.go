package main

import (
	"fmt"

	"os"

	"example.com/greetings"
)

func main() {
	name := "Daniel"

	if len(os.Args) > 1 && os.Args[1] != "" {
		name = os.Args[1]
	}

	message := greetings.Hello(name)

	fmt.Println(message)
}
