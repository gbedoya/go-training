package main

import (
	"fmt"
	"os/exec"
)

func main(){
	// Go requires the absolute path ot the binary we want to execute,
	// so we'll use exec.LookPath to find it
	//binary, lookErr := exec.LookPath("ls")
	//if lookErr != nil {
	//	panic(lookErr)
	//}

	// Exec requires arguments in slice form (as apposed to one big-string).
	//args := []string{"ls","-a", "-l", "-h"}

	// Exec also needs a set of environment variables to use.
	//env := os.Environ()

	//execErr := syscall.Exec(binary, args, env)
	//if execErr != nil {
	//	panic(execErr)
	//}

	c := exec.Command("cmd", "/K", "dir", "C:")
	if err := c.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}