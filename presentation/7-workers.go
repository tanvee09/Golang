package main

import (
	"fmt"
)

func square(num int) int {
	return num * num
}

// Unidirectional channels

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- square(n)
	}
}

func main() {
	jobs := make(chan int, 25)
	results := make(chan int, 25)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 25; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 25; j++ {
		fmt.Println(<- results)
	}
}