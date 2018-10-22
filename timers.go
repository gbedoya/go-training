package main

import (
	"fmt"
	"time"
)

func Timer() {
	// Timers represent a single event in the future.
	// You tell the timer how long you want to wait,
	// and it provides a channel that will be notified at that time.
	// This timer will wait 2 seconds.

	p = fmt.Println
	timer1 := time.NewTimer(2 * time.Second)

	// Time <- timer1.c blocks on the timer's channel C until it sends a value
	// indicating that the time expired
	<-timer1.C
	p("Time 1 Expired")

	// One reason a timer may be useful is that you can cancel the timer before it expires.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		p("Time 2 Expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		p("Time 2 Stopped")
	}

	// The first timer will expire ~2s after we start the program,
	// but th second should be stopped before it has a chance to expire
}
