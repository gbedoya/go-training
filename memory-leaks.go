package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var s0 string
var s2 []int

func substringLeak(s1 string) {
	s0 = s1[:50]
	// Now, s0 shares the same underlying memory block
	// with s1. Although s1 is not alive now, but s0
	// is still alive, so the memory block they share
	// couldn't be collected, though there are only 50
	// bytes used in the block and all other bytes in
	// the block become unavailable.

	// Remediation:
	var b strings.Builder
	b.Grow(50)
	b.WriteString(s1[:50])
	s0 = b.String()
}

func subslicesLeak(s3 []int) {
	// Assume the length of s1 is much larger than 30
	s2 = s3[len(s3)-30:]

	// Remediation
	s2 = append([]int(nil), s3[len(s3)-30:]...)
	// Now, the memory block hosting the elements
	// of s1 can be collected if no other values
	// are referencing the memory block.
}

func NotResettingPoints() []*int {
	s := []*int{new(int, new(int), new(int), new(int)}
	// do something with s ...

	return s[1:3:3]
	// the memory block allocated for the first and
	// the last elements of slice s will get lost. As long
	// as the returned slice is still alive, it will
	// prevent any element of s from being collected.

	// Remediation
	s[0], s[len(s)-1] = nil, nil
	return s[1:3:3]
	// reset the pointers stored in the lost elements

}

func finalizerLeak() {
	type T struct {
		v [1<<20]int
		t *T
	}

	var finalizer = func(t *T) {
		fmt.Println("finalizer called")
	}

	var x, y T

	// The SetFinalizer call makes x escape to heap.
	runtime.SetFinalizer(&x, finalizer)

	// The following line forms a cyclic reference
	// group with two members, x and y.
	// This causes x and y are not collectable.
	x.t, y.t = &y, &x // y also escapes to heap.
}
// time.Timer Leaking
// When a time.Timer value is not used any more,
// it will be garbage collected after some time.
// This is not true for a time.Ticker value.
// Stop a time.Ticker value when it is not used any more.

// Go runtime will not kill hanging goroutines,
// so the resources allocated for
// (and the memory blocks referenced by)
// the hanging goroutines will never get garbage collected.

// A very large deferred call stack may also consume much memory,
// and the unexecuted deferred calls may prevent some resources
// from being released in time. For example, if there are many
// files needed to be handled in a call to the following function,
// then a large number of file handlers will be not get released
// before the function exits.
func writeManyFiles(files []File) error {
	for _, file := range files {
		f, err := os.Open(file.path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(file.content)
		if err != nil {
			return err
		}

		err = f.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}

// We can use an anonymous function to enclose the deferred
// calls so that the deferred function calls will get executed earlier.
func writeManyFiles(files []File) error {
	for _, file := range files {
		if err := func() error {
			f, err := os.Open(file.path)
			if err != nil {
				return err
			}
			// The close method will be called at
			// the end of the current loop step.
			defer f.Close()

			_, err = f.WriteString(file.content)
			if err != nil {
				return err
			}

			return f.Sync()
		}(); err != nil {
			return err
		}
	}

	return nil
}

func MemoryLeaks() {

	//s := createStringWithLengthOnHeap(1 << 20) // 1M bytes
	//substringsLeak(s)

}
