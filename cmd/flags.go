package main

import (
	"flag"
	"fmt"
)

/*
go build -o flags.exe flags.go
flags.exe --word bar --num 0 --fork foo bar
*/

func main() {

	// Basic flag declarations are available for string,
	// integer, and boolean options.
	// Here we declare a string flag word with a default value "foo" and a short description.
	// This flag.String function returns a string pointer (not a string value)

	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("num",  42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")


	// It's also possible to declare an option that uses an existing
	// var declared elsewhere in the program
	// Note that we need to pass in a pointer to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared,
	// call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)			// ?
	fmt.Println("tail:", flag.Args())

}


