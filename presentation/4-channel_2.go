/*
	Receiving multiple values
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

	close(channel)
}

func main() {
	channel := make(chan string)

	go say("Hello", channel)

	// for {
	// 	msg, isOpen := <- channel
	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	for msg := range channel {
		fmt.Println(msg)
	}
}