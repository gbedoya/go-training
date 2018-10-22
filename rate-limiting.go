package main

import (
	"fmt"
	"time"
)

func RateLimit() {

	requests := make(chan int, 5)

	for i := 1; i <=5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value very 200 milliseconds
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the limiter channel before serving each request,
	// we limit ourselves to 1 request every 200 milliseconds
	for req := range requests {
		<- limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 milliseconds we'll try to add a new value to burtyLimiter,
	// up to its limit of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Simulate 5 incoming requests.
	// The first 3 of these will benefit from the burp capability of burstyLimiter
	burstyRequests := make(chan int, 5)
	for i:= 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}