package main

import "fmt"

// [1 to n^2]  ==>>  n^2 x (n^2 + 1)/2

// expSum + a - b = ActualSum
// b = expSum + a - ActualSum
// expectedSum = n^2 x (n^2 + 1) / 2  =>  (9 x 10) / 2 => 45

func missingAndRpeat(mat [][]int) []int { // Time Complexity O(n^2)
	m := map[int]int{} // Hashing DS
	var a int          // For repeating values
	var b int          // For missing values
	n := len(mat)

	ans := []int{}
	actualSum := 0

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			num := mat[i][j]
			_, exist := m[num]

			actualSum += mat[i][j]

			if exist {
				a = num
				ans = append(ans, a)
			}

			m[num] = j
		}
	}

	expectedSum := (n * n) * (n*n + 1) / 2
	b = expectedSum + a - actualSum
	ans = append(ans, b)

	return ans
}

func main() {
	// n x n grid
	grid := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	// Output := [9,5]
	// a => repeat value in grid
	// b => missing value in grid

	vec := missingAndRpeat(grid)
	fmt.Println("The repeating and missing values are :- ", vec)
}
