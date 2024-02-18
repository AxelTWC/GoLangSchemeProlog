package main

import "fmt"

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i // send i to the channel
	}
	close(ch) // close the channel when finished
}

func consumer(ch <-chan int) {
	for {
		val, ok := <-ch // receive a value from the channel
		if !ok {
			break // exit loop if the channel is closed
		}
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int) // create an unbuffered channel

	go producer(ch) // start the producer goroutine
	go consumer(ch) // start the consumer goroutine

	// wait for both goroutines to finish
	var input string
	fmt.Scanln(&input)
}

//Output
//1 2 3 4 5 (End with User Input)
