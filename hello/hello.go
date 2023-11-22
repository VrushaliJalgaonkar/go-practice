package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}

//package -  is a way to group functions, and its made up of all the files in the same directory
//to run the file -> go run .
//Visit pkg.go.dev and search for a "quote" package.
//rsc.io/quote/v4 -> This package collects pithy sayings.
//func Glass() string -> Glass returns a useful phrase for world travelers.
//func Go() string -> Go returns a Go proverb.
//func Hello() string -> Hello returns a greeting.
//func Opt() string -> Opt returns an optimization truth.

//Go code is grouped into packages, and packages are grouped into modules
