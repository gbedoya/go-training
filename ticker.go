package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future.
// Tickers are for when you want to do something repeatedly at regular intervals.
// Here's an example of a ticket that ticks periodically until we stop it.
func Ticker() {

	p = fmt.Println

	ticker := time.NewTicker(500 * time.Millisecond)
	go func(){
		for t := range ticker.C {
			p("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	p("Ticket stopped")

	// When we run this program the ticket should tick 3 times before we stop it.

}