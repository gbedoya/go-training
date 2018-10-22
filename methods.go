package main

import "fmt"

type rectangle struct {
	width, height int
}

// This area method has a receiver type of *rectangle
func (r *rectangle) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here's an example of a value receiver.
func (r rectangle) perim() int {
	return 2*r.width + 2*r.height
}

func Methods() {
	r := rectangle{10, 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
	// Go automatically handles conversion between values and pointers for methods calls.
	// You may want to use a pointers receiver type to avoid copying on method calls or
	// to allow the method to mutate the receiving struct.

}
