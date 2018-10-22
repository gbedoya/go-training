package main

import (
	"fmt"
	"time"
)

/*
Channels are the pipes the connect concurrect goroutines.
You can send values into channels and receive those values into another  goroutine.
*/
func Channels() {
	// Create a new channel with make(chan val-type)
	// Channels are typed by the values they convey
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax.
	go func() {
		messages <- "ping"
	}()

	// The <-channel syntax receives a value from the channel.
	// Here we'll receive the "ping" message we sent above and print it out.
	//msg := <- messages
	//fmt.Println(msg)

	go func() {
		m := <- messages
		fmt.Println(m)
	}()
	time.Sleep(2 * time.Second)
}
