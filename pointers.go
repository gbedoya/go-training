package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Pointers() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The &i syntax gives the memory address of i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too.
	fmt.Println("Pointer:", &i)

	// zeroval doesn't change the i in main.
	// but zeroptr does because it has a reference to the memory address

}
