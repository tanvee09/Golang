/*
	Waitgroup waits for a collection of goroutines to finish
	Add(), Done(), Wait()
*/

package main

import (
	"fmt"
	"time"
	"sync"
)

func say(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(500 * time.Millisecond)
	}
}

func sayWithPointer(text string, wg *sync.WaitGroup) {
	say(text)
	(*wg).Done()
}

func main ()  {
	var wg sync.WaitGroup
	wg.Add(1) // Increment counter
	// go sayWithPointer("Hello", &wg) // Pass WaitGroup as pointer

	// Anonymous Function
	go func () {
		say("Hello")
		wg.Done()
	}()

	wg.Wait()
}
