package main

import (
	"strconv"
	"fmt"
)

func NumberParse() {

	// With ParseFloat, this 64 tells how many bits of precision to parse.
	f, _ := strconv.ParseFloat("1.123", 64)
	fmt.Println(f)

	// For ParseInt, the 0 means infer the base from the string
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
