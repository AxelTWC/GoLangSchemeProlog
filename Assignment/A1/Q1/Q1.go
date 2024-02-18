package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type Point struct {
	x float64
	y float64
}

func main() {
	points := []Point{{8., 1.}, {3., 2.}, {7., 4.}, {6., 3.}}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if j < i {
					j++
				}
				if i == j {
					j++
				}
				if j > 3 {
					j = i + 1
					i++
				}
				if i == 3 {
					break
				}
				if j < i {
					j++
				}
				if i == j {
					j++
				}
				if j < i {
					j++
				}
				fmt.Printf("Index: %v, %v\n", i, j)
				go MidPoint(points[i], points[j])
				time.Sleep(1 * time.Second)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func MidPoint(input1 Point, input2 Point) {
	fmt.Printf("Points= (%v,%v) (%v,%v)\n", input1.x, input1.y, input2.x, input2.y)
	midpointx := (input1.x + input2.x) / 2
	midpointy := (input1.y + input2.y) / 2
	Midpoint := []Point{{midpointx, midpointy}}
	fmt.Printf("Midpoint = (%v, %v)\n", Midpoint[0].x, Midpoint[0].y)
	distance := math.Sqrt((math.Pow(input2.x-input1.x, 2)) + (math.Pow(input2.y-input1.y, 2)))
	fmt.Printf("Length = %v\n", math.Ceil(distance*100)/100)
}
