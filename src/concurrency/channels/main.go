package main

import "fmt"

func main() {
	// Initialize a channel and the data that's passed through this channel
	myChannel := make(chan string)

	// go routine
	// anonymous function
	go func() {
		// putting the data onto the channel
		myChannel <- "data"
	}()

	// this will actually wait for either this channel to close
	// or for a massage to be received
	msg := <-myChannel

	fmt.Println(msg)
}
