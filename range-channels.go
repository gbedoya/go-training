package main

import "fmt"

func RangeChannels(){

	// We'll iterate over 2 values in the queue channel.
	 queue := make(chan string, 2)
	 queue <- "one"
	 queue <- "two"
	 close(queue)

	 for elem := range queue {
	 	fmt.Println(elem)
	 }
}
