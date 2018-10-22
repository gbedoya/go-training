package main

import (
	"fmt"
	"os"
)

/*
Command-line arguments are a common way to parameterize execution of programs.
For example, go run hello.go uses run and hello.go to the go program.
go build -o args.exe args.go
args.exe 1 2 3 4
*/

func main() {

	print := fmt.Println

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	print(argsWithProg)
	print(argsWithoutProg)
	print(arg)


}