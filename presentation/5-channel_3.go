/*
	Buffered channels to avoid deadlocks

	Sending to a channel blocked only when buffer is full
	Receiving from a channel blocked only when buffer is empty
*/

package main

import (
	"fmt"
)

func main() {
	channel := make(chan string, 2)
	channel <- "Hello"
	channel <-"Hello" 
	
	msg := <-channel
	fmt.Println(msg)
}