package main

import (
	"fmt"
	"math"
)

// Maximum Subarray ->> Continuous part of the given array / Possible Sub-arrays from the given array
// [1,2,3,4,5]
// 1,2,3,4,5
// 2,3,4,5,
// 3,4,5
// 4,5
// 5

func main() {
	var n = 5
	var arr = [5]int{1, 2, 3, 4, 5}

	fmt.Println("----------------- Possible Sub-Arrays from the given Array ---------------------")
	for st := 0; st < n; st++ {
		for end := st; end < n; end++ {
			for i := st; i <= end; i++ {
				fmt.Print(arr[i])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}

	fmt.Println("----------------- Maximum Sub-Array Sum from the given Array ---------------------")
	// Brute Force Approach  -->>  Take out maximum subarray and then sum all of them
	arr1 := [7]int{3, -4, 5, 4, -1, 7, -8}
	sz := 7

	maxSum := math.MinInt
	for st := 0; st < sz; st++ {
		currSum := 0
		for end := st; end < sz; end++ {
			// fmt.Print(arr1[end], " ")
			currSum += arr1[end]
			maxSum = max(currSum, maxSum)
		}
	}

	fmt.Print("The maximum sub-array sum from the given array :- ", maxSum)
}
