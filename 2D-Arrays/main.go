// In programming, the 2D - Array can also be called as Matrix or Matrices

// Data is stored in the form of table like rows or columns
// 2D-Array are two diemensional array

// matrix = [no. of rows] [no. of columns]
// matrix := [4][3]  // 2D-Array

package main

import (
	"fmt"
)

func matrixElement(matrix [4][3]int, rows, cols int) {

	// Outer Loop will control our rows
	// Inner Loop will control our columns

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	// Initialization of 2D Matrix
	matrix := [4][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	rows := 4
	cols := 3

	// fmt.Println(matrix[2][1]) // 8
	// fmt.Println(matrix[2][2]) // 9

	// Print all the matrix element
	matrixElement(matrix, rows, cols)
}

// 2D-Arrays are stored in memory in two ways :-
// . This will depend on the system

// 1. Row Major :- Linear storage
// [1,2,3,4,5,6,7,8,9]

// 2. Column Major : Stored like a column wise fashion
// [1,4,7, 2,5,8, 3,6,7]
