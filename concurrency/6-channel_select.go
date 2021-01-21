/*
	Select reads from whichever channel is ready to send
	Picks one at random if more than one channel is ready to send
*/

package main

import (
	"fmt"
	"time"
)


func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Sending every 500 ms"
			time.Sleep(500 * time.Millisecond)
		}
		
	}()

	go func() {
		for {
			c2 <- "Sending every 2000 ms"
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	for {
		select {
			case msg1 := <-c1 :
				fmt.Println(msg1)
			case msg2 := <-c2 :
				fmt.Println(msg2)
		}
	}
}