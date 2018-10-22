package main

import "fmt"

func Variables() {
	// Declare 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// Declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b,c)

	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// Variables declared without a initialization are zero-valued
	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}