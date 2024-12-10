package main

import "fmt"

func spiralMatrix(mat [][]int) []int {
	m := len(mat)    // no. of rows
	n := len(mat[0]) // no. of columns

	srow := 0 // Starting row
	scol := 0 // Starting column

	erow := m - 1
	ecol := n - 1

	var vector []int

	for srow <= erow && scol <= ecol { // Time Complexity = O(m x n) => O(n)
		// Top Boundary
		for j := scol; j <= ecol; j++ {

			// if len(vector) <= srow {
			// 	// Ensure vector has a row corresponding to srow
			// 	vector = append(vector, []int{})
			// }

			// Append element to the current row
			vector = append(vector, mat[srow][j])

		}

		// Right Boundary
		for j := srow + 1; j <= erow; j++ {

			// if len(vector) <= j {
			// 	// ensure vector has a corresponding right column
			// 	vector = append(vector, []int{})
			// }

			vector = append(vector, mat[j][ecol])
		}

		// Edge Case will be noted for bottom and left boundary in which the bottom boundary will interect with top boundary and left will intersect with right boundary

		// Bottom Boundary
		for j := ecol - 1; j >= scol; j-- {
			if srow == erow { // to stop the repition of element
				break
			}

			// if len(vector) <= erow {
			// 	vector = append(vector, []int{})
			// }

			vector = append(vector, mat[erow][j])
		}

		// Left Boundary
		for j := erow - 1; j >= srow+1; j-- {
			if scol == ecol { // to stop the repition of element
				break
			}

			// if len(vector) <= j {
			// 	vector = append(vector, []int{})
			// }

			vector = append(vector, mat[j][scol])
		}

		srow++
		erow--
		scol++
		ecol--
	}

	return vector

}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	fmt.Println(spiralMatrix(matrix))
}
