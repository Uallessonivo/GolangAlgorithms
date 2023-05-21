package main

import "fmt"

func main() {
	// Initialize a channel and the data that's passed through this channel
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	// go routine
	// anonymous function
	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "another data"
	}()

	// Executes when receives a msg from one channel or another
	// if both channels have message this will select a random channel
	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
}
