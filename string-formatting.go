package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func StringFormatting() {

	var pf = fmt.Printf

	// Go offers several printing "verbs" designed to format general Go values.
	// For example, this prints an instance of our point struct.
	p := point{1, 2}
	pf("%v\n", p)

	// The %+v variant will include the struct's field names.
	pf("%+v\n", p)

	// The %#v variants prints a Go syntax representation of the value,
	// i.e. the source code snippet
	pf("%#v\n", p)

	// To print the type of a value, use %T
	pf("%T\n", p)

	// base-10 formatting
	pf("%d\n", 123)

	// Print a binary representation
	pf("%b\n", 7)

	// Prints the character corresponding to the given integer
	pf("%c\n", 65)

	// Formatting booleans is straight-forward
	pf("%t\n", true)

	// for hex encoding
	pf("%x\n", "f")

	// for floats
	pf("%f\n",78.9)

	// %e and %E format the float in (slightly different versions of) scientific notation
	pf("%E\n", 1234000000.0)
	pf("%e\n", 1234000000.0)

	// For string printing use %s
	pf("%s\n", "\"string\"")

	// To double-quote strings as in Go source, use %q
	pf("%q\n", "\"string\"")

	// To print a representation of a pointer, use %p
	pf("%p\n", &p)

	// To specify the width of an integer, use a number after the % in the verb.
	// By default the result will be right-justified and padded with spaces.
	pf("|%6d|%6d|\n", 12, 345)

	// width.precision syntax
	pf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// You may also want to control width when formatting strings
	pf("|%6s|%6s|\n", "foo", "bar")

	// TO left-justify use the - flag as with numbers.
	pf("|%-6s|%-6s|\n", "foo", "bar")

	// Printf formats strings to os.Stdout
	// Sprintf formats and returns a string w/out printing
	s := fmt.Sprintf("a %q", "string")
	fmt.Println(s)

	// You can format+print to is.Writers other than os.Stdout using Fprintf
	fmt.Fprintf(os.Stderr, "an %q", "error")

}
