package main

import "fmt"

type person struct {
	name string
	age int
}

func Struct() {

	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{"Ann", 40})

	// Access struct fields with a dot.
	s := person{"Sean", 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers
	// The pointers are automatically de-referenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}

