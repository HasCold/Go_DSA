package main

import (
	"fmt"
	"math"
)

// Outer loop will control the no. of rows or lines
// Inner loop will control the no. of columns or the data present inside the lines

func maxColSum(mat [3][3]int) (int, map[string]int) {

	curSum := math.MinInt
	pair := map[string]int{"col": -1} // [key_type]value_type

	for j := 0; j < len(mat); j++ {
		columnSum := 0

		for i := 0; i < len(mat[j]); i++ {
			columnSum += mat[i][j]
		}

		if curSum < columnSum {
			pair["col"] = j
			curSum = columnSum
		}
	}

	return curSum, pair
}

func main() {
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	MaxSum, Pair := maxColSum(matrix)
	fmt.Printf("The Maximum Column Sum from 2D Matrix is %v and their available pair of column is %v", MaxSum, Pair)
}

// Format 	Specifier							Meaning	Example
// %c		Character(from ASCII/Unicode)		fmt.Printf("%c", 65) → A
// %d		Decimal integer						fmt.Printf("%d", 123) → 123
// %v		Default format (any type)			fmt.Printf("%v", 123) → 123, fmt.Printf("%v", "Hello") → Hello
