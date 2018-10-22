package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func Constants () {
	fmt.Println(s)

	// A const statement can appear anywhere a var statement can.
	const n = 500000000
	// Constant expressions perform arithmetic with arbitrary precision.
	const d = 3e10 / n

	// A numeric constant has no type until it's given one, such as by an explicit cast
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))


}