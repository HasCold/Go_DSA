package main

import "fmt"

func missingAndRpeat(mat [][]int) (int, int) {
	m := map[int]int{} // Hashing DS
	var a int          // For repeating values
	var b int          // For missing values

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			num := mat[i][j]
			_, exist := m[num]

			if exist {
				a = num
				break
			}

			m[num] = j
		}
	}

	return a, b
}

func main() {
	// n x n grid
	grid := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	// Output := [9,5]
	// a => repeat value in grid
	// b => missing value in grid

	repeat, miss := missingAndRpeat(grid)
	fmt.Println("The repeating and missing values are :- ", repeat, miss)
}
