package main

// Practicing concurrency patterns

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	// Will be executed but not necessary in this order
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	// This is used only for await the return form the async executed functions
	time.Sleep(time.Second * 2)

	fmt.Println("hi")
}
