package main

import (
	"fmt"
	"sync"
)

func fct(line []float64) {
	for _, v := range line {
		fmt.Printf("%f ", v)
	}
}

func fct2(matrix [][]float64) {
	matrix[2][0] = 12345.6
}

func sort(matrix []float64) []float64 {
	for i := 0; i <= len(matrix)-1; i++ {
		for j := 0; j < len(matrix)-1-i; j++ {
			if matrix[j] > matrix[j+1] {
				matrix[j], matrix[j+1] = matrix[j+1], matrix[j]
			}
		}
	}
	return matrix
	// fmt.Print(matrix)
}
func transpose(tab [][]float64) [][]float64 {
	n := len(tab)
	m := len(tab[0])
	transposed := make([][]float64, m)
	for i := range transposed {
		transposed[i] = make([]float64, n)
		for j := range transposed[i] {
			transposed[i][j] = tab[j][i]
		}
	}
	return transposed

}
func sortSeperateThread(wg *sync.WaitGroup, currentRow []float64, res chan []float64) {
	defer wg.Done()
	currentSortedRow := sort(currentRow)
	// fmt.Print(currentRow)
	res <- currentSortedRow
}
func sortRows(tab [][]float64) [][]float64 {
	// n := len(tab)
	// m := len(tab[0])
	// sortedRows := make([][]float64, n)
	// for i := range sortedRows {
	// 	sort(tab[i])
	// 	sortedRows[i] = make([]float64, m)
	// 	for j := range sortedRows[i] {
	// 		sortedRows[i][j] = tab[i][j]
	// 	}
	// }
	// return sortedRows
	n := len(tab)
	// m := len(tab[0])
	sortedRows := make([][]float64, n)
	i := 0
	var wg sync.WaitGroup
	sortedRows = tab
	for _, needToSort := range sortedRows {
		wg.Add(1)
		res := make(chan []float64)
		// fmt.Print(needToSort)
		go sortSeperateThread(&wg, needToSort, res)
		temp := <-res
		sortedRows[i] = temp
		// fmt.Print(sortedRows[i])
		for j := range sortedRows[i] {
			sortedRows[i][j] = tab[i][j]
		}
		i++
	}
	wg.Wait()
	// fmt.Print(sortedRows)
	return sortedRows
}
func main() {
	// array := [][]float64{{7.1, 2.3, 1.1},
	// {4.3, 5.6, 6.8},
	// {2.3, 2.7, 3.5},
	// {4.5, 8.1, 6.6}}
	array := [][]float64{{1.1, 7.3, 3.2, 0.3, 3.1},
		{4.3, 5.6, 1.8, 5.3, 3.1},
		{1.3, 2.7, 3.5, 9.3, 1.1},
		{7.5, 5.1, 0.6, 2.3, 3.9}}
	fmt.Printf("Initial Array:")
	fmt.Printf("\n")
	fmt.Printf("[")
	fct(array[0])
	fmt.Printf("]")
	fmt.Printf("\n")
	fmt.Printf("[")
	fct(array[1])
	fmt.Printf("]")
	fmt.Printf("\n")
	fmt.Printf("[")
	fct(array[2])
	fmt.Printf("]")
	fmt.Printf("\n")
	fmt.Printf("[")
	fct(array[3])
	fmt.Printf("]")
	// fmt.Print(sortRows(array))
	// fmt.Print(transpose(array))
	// fmt.Print(sortRows(transpose(array)))
	fmt.Printf("\n")
	fmt.Printf("Final result:")
	fmt.Printf("\n")
	fmt.Print(transpose(sortRows(transpose(array))))

	// fct2(array)
	// fct(array[2][:])
	// fmt.Print(sort(array[2]))
	// fmt.Print(transpose(array))
	// fmt.Print(sortRows(array))

}
