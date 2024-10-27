// Description: This program demonstrates how to use channels to pass values between goroutines.
// Based on concurrency example from: https://fireship.io/lessons/learn-go-in-100-lines/

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // Create a channel to pass ints
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c) // Start a goroutine
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c // Receive a value from a channel
		fmt.Println("gopher", gopherID, "finished the dish")
	} // All goroutines are finished at this point
}

/* Notice the channel as an argument */
func cookingGopher(id int, c chan int) {
	fmt.Println("gopher", id, "started cooking")
	time.Sleep(1 * time.Second) // Added a delay to show routine execution
	c <- id // Send a value back to main
}