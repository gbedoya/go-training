package main

import (
	"os"
	"fmt"
)

/*
Defer is used to ensure that a function call is performed later in a program's execution,
usually for purposes of cleanup.
defer is often used where e.g. ensure and finally would be used in other languages.
 */

func Defer() {
	//f := createFile("/tmp/defer.txt")
	f := createFile("C:/Users/geller/go/src/go-training/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("Creating ...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing ...")
	fmt.Fprintf(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing ...")
	f.Close()
}
