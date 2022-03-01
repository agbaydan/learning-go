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

	names := []string{}
	for i, name := range os.Args {
		// skip first arg, this is a very lazy way to do this lol
		if i == 0 {continue}

		names = append(names, name)
	}

	greetings, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greetings)
}
