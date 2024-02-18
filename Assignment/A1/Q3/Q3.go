package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func RandomGenerator(wg *sync.WaitGroup, stop <-chan bool, m int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer func() {
			wg.Done()
		}()
		defer close(intStream)
		for {
			select {
			case <-stop:
				return
			case intStream <- m * rand.Intn(1000000/m):
			}
		}
	}()
	return intStream
}
func Multiple(x int, m int) bool {
	return x%m == 0
}
func main() {
	wg := &sync.WaitGroup{}
	stop := make(chan bool)
	defer close(stop)
	generatorMultipleOf5 := RandomGenerator(wg, stop, 5)
	generatorMultipleOf13 := RandomGenerator(wg, stop, 13)
	generatorMultipleOf97 := RandomGenerator(wg, stop, 97)
	var count5, count13, count97 int
	for i := 0; i < 100; i++ {
		select {
		case x := <-generatorMultipleOf5:
			if Multiple(x, 5) {
				count5++
			}
		case x := <-generatorMultipleOf13:
			if Multiple(x, 13) {
				count13++
			}
		case x := <-generatorMultipleOf97:
			if Multiple(x, 97) {
				count97++
			}
		}
	}
	fmt.Printf("Total multiples of 5: %d\n", count5)
	fmt.Printf("Total multiples of 13: %d\n", count13)
	fmt.Printf("Total multiples of 97: %d\n", count97)
	wg.Wait()
}
