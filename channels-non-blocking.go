package main

import "fmt"

func NonBlockingChannels(){
	print := fmt.Println

	messages := make(chan string)
	signals := make(chan bool)

	// Here's a non-blocking receive.
	// If a value is available on messages then select will take the <-messages case with the value.
	// If not it will immediately take the default case.
	select {
	case msg := <- messages:
		print("received:", msg)
	default:
		print("No message received.")
	}

	// A non-blocking send works similarly.
	// Here msg cannot be sent to the message channel,
	// because the channel has no buffer and there is no receiver.
	// Therefore the default case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		print("sent message.")
	default:
		print("no message sent")
	}

	// We can use multiple cases above the default clause to implement a multi-way non-blocking select.
	// Here we attempt non-blocking receives on both messages and signals.
	select {
	case msg := <-messages:
		print("sent message", msg)
	case sig := <-signals:
		print("received signal", sig)
	default:
		print("no activity")
	}
}
