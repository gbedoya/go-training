package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures

// Here's a basic interface for geometric shapes.

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interfaces() {
	r := rect{3, 4 }
	c := circle{ 5}

	// The circle and rect struct types both implement the geometry interaces
	// so  we can use instances of these structs as arguments to measure.
	measure(r)
	measure(c)
}
