// You are given an m x n integer matrix matrix with the following two properties:

// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.

// You must write a solution in O(log(m * n)) time complexity.

// If in the question will demand for the logarithmic time complexity O(log n) then we must think about the Binary Search Algo
// Binary Search Algorithm can only be implement on the sorted array

package main

import (
	"fmt"
)

func search2DMatrix(mat [][]int, tar int) bool {
	r := len(mat)    // length of row
	c := len(mat[0]) // length of column
	stRow := 0       // Starting At Row
	endRow := r - 1  // Ending At Row

	// 1. BS apply on total no. of rows
	// 2. Then BS apply on column between the range of row

	for stRow <= endRow {
		// To prevent the overflow of integer we optimize the formula
		midRow := stRow + (endRow-stRow)/2

		if mat[midRow][0] <= tar && tar <= mat[midRow][c-1] { // To check the target exist in between the range of this row
			// Search In Row
			return searchInRow(mat, tar, midRow)

		} else if tar > mat[midRow][c-1] { // To check whether the target is greater than start of the row element
			// Target is in a later row
			stRow = midRow + 1

		} else if tar < mat[midRow][0] { // To check whether the target is less than endRow of the row element
			// Target is in an earlier row
			endRow = midRow - 1
		}
	}

	return false
}

func searchInRow(matrix [][]int, target, midRow int) bool { // Time Complexity O(logn)

	n := len(matrix[0]) // no. of columns
	st := 0
	end := n - 1

	for st <= end {
		// To prevent the overflow of integer we optimize the formula
		mid := st + (end-st)/2

		if target == matrix[midRow][mid] {
			return true

		} else if matrix[midRow][mid] < target {
			st = mid + 1

		} else if matrix[midRow][mid] > target {
			end = mid - 1
		}
	}

	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}} // 2D SLices or Vector
	target := 34

	fmt.Println("The target is available in the 2D Matrix :- ", search2DMatrix(matrix, target))
}
