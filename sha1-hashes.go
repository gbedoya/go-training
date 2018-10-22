package main

import (
	"crypto/sha1"
	"fmt"
)

func SHA1Hashes() {

	s := "sha1 this string"
	h := sha1.New()
	h.Write([]byte(s))
	// The argument to Sum can be used to append to an existing byte slice
	bs := h.Sum(nil)

	// Use the %x format verb to convert a hash results to a hex string
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
