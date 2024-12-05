// In 1-diemensional Array we let the index (i++) for position of an element
// In 2-diemensional Array we suppose cell instead of index ; index -> (i,j)

package main

import "fmt"

// Outer loop tells you the no. of lines or no. of rows in a cell.
// Inner loop tells you the no. of columns in a cell or the thing that you want to print.

func linearSearchOn2DArray(matrix [4][3]int, find int) bool { // we must specify the columns as compiler needs this info
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == find {
				return true
			}
		}
	}

	return false
}

func main() {

	matrix := [4][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	find := 8

	fmt.Println(len(matrix))
	fmt.Println("Number is available :- ", linearSearchOn2DArray(matrix, find))
}

// [1,2,3]
// [4,5,6]
// [7,8,9]
