package main

import (
	"fmt"
)

// If the rows and columns are equal means n = 3(odd number) then it must have one common element
// If the rows and columns are equal means n = 4(even number) then there is no common element

// Primary Diagonal :-  i = j ; [0][0]
// Secondary Diagonal :-  j = n-1-i ; [0][3]

// Outer loop will control the no. of rows or lines (i)
// Inner loop will control the no. of columns or the data present inside the lines (j)

func diagonalMatrix(mat [4][4]int, n int) int {

	// pdSum := 0 // Primary Diagonal
	// sdSum := 0 // Secondary Diagonal
	sum := 0

	for i := 0; i < n; i++ {
		// Brute Approach
		// for j := 0; j < n; j++ {  // O(n^2)
		// 	if i == j {
		// 		pdSum += mat[i][j]

		// 	} else if j == n-1-i {
		// 		sdSum += mat[i][j]

		// 	}
		// }

		// Optimized Approach
		for i := 0; i < n; i++ { // O(n)
			sum += mat[i][i] // if i==j for pdSum

			if i != n-i-1 { // To avoid the common values
				sum += mat[i][n-1-i]
			}
		}
	}

	return sum
}

func main() {
	matrix := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	n := 4 // length of matrix

	// 2D vector or slices for static arrays
	mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	diagSum := diagonalMatrix(matrix, n)
	fmt.Println(diagSum, mat)
}
