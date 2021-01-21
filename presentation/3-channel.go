/*
	Channels are like pipes through which you can send and recieve values
	Channels are pointers by default

	channel := make(chan string)
	channel <- msg
	msg := <-channel
*/

package main

import (
	"fmt"
	"time"
)

func say(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// var channel chan string // Will create null channel
	
	channel := make(chan string)
	
	go say("Hello", channel)

	msg := <-channel
	fmt.Println(msg)

	// time.Sleep(10000 * time.Millisecond)
}

/*
	Goroutine 1		Goroutine 2
	1. ...			1. ...
	2. ...			2. ...
	3. ...			3. ch <- "Hello"
	4. ...			4. ...
	5. ...			5. ...
	6. msg := <- c	6. ...
	7. ...			7. ...
*/