package main

import (
	"fmt"

	"log"

	"os"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined logger, included the log entry prefix
	// and a flag to disable printing the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	name := ""
	if len(os.Args) > 1 && os.Args[1] != ""{
		name = os.Args[1]
	}

	message, err := greetings.Hello(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
