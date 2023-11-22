package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Vrushali", "Ritik", "Rahul", "John"}
	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(messages)
}

//go mod edit -replace example.com/greetings=../greetings --> to adapt the example.com/hello module
//so it can find the example.com/greetings code on your local file system.
//go mod tidy --> so that example.com/hello downloads example.com/greetings from local file system
//Compile & install application --> https://go.dev/doc/tutorial/compile-install
