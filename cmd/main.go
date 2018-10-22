package main

import (
	"os/exec"
	"fmt"
)
// go build -o spawn.exe main.go
// go run main.go

func main() {

	dateCmd := exec.Command("cmd", "/K", "date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> Date")
	fmt.Println(string(dateOut))
}