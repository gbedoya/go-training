package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func Regex() {

	p := fmt.Println

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	p(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	p(r.MatchString("peach"))

	p(r.FindString("peach punch"))

	// Finds the first match but returns the start
	// and end indexes for the match instead of the matching text
	p(r.FindStringIndex("peach punch"))

	// Will return information for both p([a-z]+ch and (a-z)+
	p(r.FindStringSubmatch("peach punch"))
	p(r.FindStringSubmatchIndex("peach punch"))

	// The All variants of these functions apply to all matches in the input,
	// not just the first. For example to find all matches for a regexp.
	p(r.FindAllString("peach punch pinch azerty", -1))
	p(r.FindAllStringIndex("peach punch pinch azerty", -1))
	p(r.FindAllStringSubmatchIndex("peach punch pinch azerty", -1))

	// Providing a non-negative integer will limit the number of matches
	p(r.FindAllString("peach punch pinch azerty", 2))

	p(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	p(r)

	p(r.ReplaceAllString("peach punch pinch azerty", "<replace>"))

	// The Func variant allows you to transform matched text with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	p(string(out))
}
