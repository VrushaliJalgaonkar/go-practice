package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	//if no mane is given return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	//returns a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!", "Great to see you, %v!", "Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

//A function whose name starts with a capital letter can be called by a function not in the same package.
//This is known in Go as an exported name.
//:= operator is a shortcut for declaring and initializing a variable in one line
//randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
//In Go, you initialize a map with the following syntax: make(map[key-type]value-type).
//for loop, range returns two values: the index of the current item in the loop and a copy of the item's value.
//You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.
