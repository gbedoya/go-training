package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// The primary mechanism for managing state in Go is communication over channels.
// We saw this for example with workers pools.
// There are few other options for managing state though.
// Here we'll look at using sync/atomic package for atomic counters accessed by multiple goroutines

func AtomicCounter() {
	// We'll use an unsigned integer to represent our (always-positive) counter.
	var ops uint64

	// To simulate concurrent updates,
	// we'll start 50 goroutines that each increment the counter about once a millisecond.
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// To atomically increment the counter we use AddUint64,
				// giving it the memory address of our ops with the & syntax.
				atomic.AddUint64(&ops, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}
	// Wait a second to allow some ops to accumulate
	time.Sleep(5 * time.Second)

	// In order to safely use the counter while it's still being updated by other goroutines,
	// we extract a copy of the current value into OpsFinal via LoadUint64.

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
