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

func maxSubArray(nums []int) int {
	maxSum := math.MinInt // -infinity
	curSum := 0
	for st := 0; st < len(nums); st++ {
		curSum += nums[st]
		maxSum = max(curSum, maxSum)
		if curSum < 0 {
			curSum = 0
		}
	}
	return maxSum
}

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

	fmt.Println("----------------- Brute-Force Approach Maximum Sub-Array Sum from the given Array ---------------------")
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
	fmt.Println("The maximum sub-array sum from the given array :- ", maxSum)

	fmt.Println("----------------- Kadane's Algorithm Maximum Sub-Array Sum from the given Array ---------------------")
	// Kadane's Algorithm works on the intuition :-
	// +ve + +ve(big number) = +ve
	// -ve + +ve(big number) = +ve
	// +ve + -ve(big number) = -ve  // bigger negative shouldn't contribute to our answer

	// we have to ignore the -ve value stored in curSum that is lower than zero and instead of that we can consider the 0 acc to Kadane's algorithm approach
	num := maxSubArray(arr1[:]) // It has linera time-complexity O(n)
	fmt.Println("Optimize num from Kadane's Algo :- ", num)
}
