package main

import (
	"fmt"
	"math"
)

func floatToInt(input float64) (floorVal int, ceilVal int) {
	floorVal = int(math.Floor(input))
	ceilVal = int(math.Ceil(input))
	fmt.Println("Input: ", input)
	return
}
func main() {
	var a, b int
	a, b = floatToInt(2.4)

	fmt.Println("Floor: ", a)
	fmt.Println("Ceil: ", b)
}
