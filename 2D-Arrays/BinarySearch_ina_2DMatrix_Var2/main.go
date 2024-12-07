// Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:

// Integers in each row are sorted in ascending from left to right.
// Integers in each column are sorted in ascending from top to bottom.

// We can apply binary search in a different variation like search apply on the column not on the row wise. And the columns will be discarded

package main

import "fmt"

func searchInMatrix(mat [][]int, tar int) bool { // O(log n)
	row := len(mat)    // No. of rows
	col := len(mat[0]) // No. of columns
	r := 0
	c := col - 1

	// We are selecting the mid as the extent side of 2D array or matrix

	for r < row && c >= 0 {
		if tar == mat[r][c] {
			return true
		} else if tar < mat[r][c] {
			c--
		} else if tar > mat[r][c] {
			r++
		}
	}

	return false
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 5

	fmt.Println("The number in 2D Matrix is :- ", searchInMatrix(matrix, target))
}
