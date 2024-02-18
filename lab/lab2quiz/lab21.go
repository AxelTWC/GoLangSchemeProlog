package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int, name string, sleep int) {
	time.Sleep(time.Second * time.Duration(sleep))

	sum := 0

	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(a[0:3], c, "A", 1)
	go sum(a[3:6], c, "B", 1)
	x := <-c
	y := <-c

	fmt.Println(x, y)
}
