package main

import(
	"fmt"
	"io/ioutil"
	"regexp"
	)

var digitRegexp = regexp.MustCompile("[0-9]+")

// https://golang.org/pkg/io/ioutil/
// FindDigits find digits in file in a consistent way
func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	c = append(c, b...)
	return digitRegexp.Find(b)
}

func Slices() {

	s := make([]string, 3)
	fmt.Println("emp:", s)
	fmt.Printf("Type: %T\n", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// Slices can also be copy'd.
	// Here we create an empty slice c of the same length as s
	// and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	fmt.Printf("Type of C: %T\n", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)
	fmt.Printf("Type of l: %T\n", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	fmt.Printf("Type of t: %T\n", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	a := []string{"Ismail", "BahaEddine"}
	b := []string{"Podba", "Mane"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	fmt.Println("append 2 slices: ", a)

	// Some gotcha
	// http://blog.golang.org/2011/01/go-slices-usage-and-internals.html
	fmt.Println("found digits: ", FindDigits("./file.txt"))
}
