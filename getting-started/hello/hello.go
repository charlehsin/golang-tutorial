package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Use other Go module - quote
	fmt.Print("Going to print a message from quote module: ")
	fmt.Println(quote.Go())

	// Use local Go module - greetings, with error handling
	// TODO: If you want to see error, pass "" instead of "Gladys".
	message, err := greetings.Hello("Gladys")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// Use local Go module - greetings, with a slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}
	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Going to print a message from local greetings module: ")
	fmt.Println(message)
	fmt.Print("Going to print a message from local greetings module: ")
	fmt.Println(messages)
}
