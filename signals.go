package main

import (
	"fmt"
	"os"
	"os/signal"
)

func Signals() {
	p := fmt.Println

	// Go signal notification works by sending
	// os.Signal values on a channel
	// We'll create a channel to receive these notifications
	// We'll als make one to notify us when the program can exit
	/*sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	signal.Notify(sigs)

	go func() {
		sig := sigs
		p()
		p(sig)
		done <- true
	}()

	p("awaiting signal")
	<-done
	p("exiting")
	*/

	c := make(chan os.Signal, 1)
	signal.Notify(c)

	// Block until a signal is received.
	s := <-c
	p("Got signal:", s)

}
