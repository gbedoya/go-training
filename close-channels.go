package main

import "fmt"

// Closing a channel indicates the no more values will be sent on it.
// This can be useful to communicate completion to the channel's receivers.


func CloseChannel() {

	print := fmt.Println

	// In this example, we'll use a jobs channel to communicate work to be done from the CloseChannel()
	// goroutine to a worker goroutine. When we have no more jobs for the worker we'll close the jobs channel.
	jobs := make(chan int, 5)
	done := make(chan bool)

	// There's the worker goroutine.
	// It repeatedly receives from jobs with j, more := < jobs. In this special 2-value from of receive,
	// the more value will be false if jobs has been closed and all values in the channel have already been received.
	// We use this to notify on done when we've worked all our jobs.
	go func() {
		for {
			j, more := <-jobs
			if more {
				print("received jobs:", j)
			} else {
				print("received all jobs")
				done <- true
				return
			}
		}
	}()


	// this sends 3 jobs to the worker over the jobs channel, then closes it.
	for j := 11; j <= 13; j++ {
		jobs <- j
		print("Sent Job:", j)
	}
	close(jobs)

	// We await the worker using the synchronization approach we saw earlier.
	<-done
}
