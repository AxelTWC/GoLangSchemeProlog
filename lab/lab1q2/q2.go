package main

import (
	"fmt"
	"strings"
)

func main() {

	lineWidth := 5
	symb := "x"
	lineSymb := symb
	formatStr := fmt.Sprintf("%%%ds\n", lineWidth)
	fmt.Printf(formatStr, lineSymb)
	for i := 2; i < lineWidth+1; i++ {
		lineSymb := strings.Repeat(symb, i)
		formatStr := fmt.Sprintf("%%%ds\n", lineWidth)
		fmt.Printf(formatStr, lineSymb)
	}
	for j := lineWidth - 1; j > 0; j-- {
		lineSymb := strings.Repeat(symb, j)
		formatStr := fmt.Sprintf("%%%ds\n", lineWidth)
		fmt.Printf(formatStr, lineSymb)
	}
}
