package main

import (
	b64 "encoding/base64"
	"fmt"
)

func Base64() {

	data := "abc123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible base64
	// THe encoder requires a []byte so we cast our string to that type.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Print(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println()
	fmt.Println(sDec)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(uDec)
	fmt.Println(string(uDec))
	// Note: trailing + vs -

}
