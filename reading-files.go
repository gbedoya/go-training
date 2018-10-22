package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkRead(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFiles() {

	// Docs: https://golang.org/pkg/io/ioutil/
	// Loads contents into memory
	dat, err := ioutil.ReadFile("C:/Users/geller/go/src/go-training/defer.txt")
	checkRead(err)
	fmt.Println(string(dat))

	// Docs: https://golang.org/pkg/os/
	f, err := os.Open("C:/Users/geller/go/src/go-training/defer.txt")
	defer f.Close()
	checkRead(err)

	// Read 5 bytes
	b1 := make([]byte, 10)
	n1, err := f.Read(b1)
	checkRead(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// Seek to a known location in the file and Read from there.
	o2, err := f.Seek(5,0)
	checkRead(err)
	b2 := make([]byte, 5)
	n2, err := f.Read(b2)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(5,0)
	checkRead(err)
	b3 := make([]byte, 5)
	n3, err := io.ReadAtLeast(f, b3, 5)
	checkRead(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//Seek(0,0) sets the file position to the beginning
	// of the file of the given stream
	_, err = f.Seek(0, 0)
	checkRead(err)

	// Docs: https://golang.org/pkg/bufio/
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(10)
	checkRead(err)
	fmt.Printf("10 bytes: %s\n", string(b4))

}
