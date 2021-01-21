/*
	Creating goroutines:
	go function_call()
*/

package main

import (
	"fmt"
	"time"
)

func say(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	go say("Hello")
	go say("Bye")

	// time.Sleep(10000 * time.Millisecond)
	fmt.Scanln()
}