package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func checkWrite(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteFiles() {

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("C:/Users/geller/go/src/go-training/dat1", d1, 0644)
	checkWrite(err)

	f, err := os.Create("C:/Users/geller/go/src/go-training/dat2")
	checkWrite(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10} // -> ASCII
	n2, err := f.Write(d2)
	checkWrite(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	checkWrite(err)
	fmt.Printf("wrote %d bytes\n", n3)
	// Issue a Sync to flush writes to stable storage.
	f.Sync()


	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	// Flush to ensure all buffered operations have been applied
	w.Flush()

}
