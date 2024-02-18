package main

import "fmt"

// Define a suitable structure
type Course struct {
	code      string
	NStudents int
	Professor string
	Avg       float64
}

func main() {
	// Create a dynamic map m
	var m = map[string]*Course{}
	m["CSI2110"] = &Course{code: "CSI2110", NStudents: 186, Professor: "Lang", Avg: 79.5}
	m["CSI2101"] = &Course{code: "CSI2101", NStudents: 211, Professor: "Moura", Avg: 81}

	// Add the courses CSI2120 and CSI2110 to the map

	for k, v := range m {
		fmt.Printf("Course Code: %s\n", k)
		fmt.Printf("Number of students: %d\n", v.NStudents)
		fmt.Printf("Professor: %s\n", v.Professor)
		fmt.Printf("Average: %f\n\n", v.Avg)
	}
}
