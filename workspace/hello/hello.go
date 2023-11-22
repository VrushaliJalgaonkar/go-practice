package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(1234))
}

//The go work init command tells go to create a go.work file for a workspace containing the modules in the ./hello directory.
