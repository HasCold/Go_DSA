// Problem Statement:
// Given an array of positive integers nums and a positive integer target, find the minimal length of a contiguous subarray of which the sum is at least target. If there is no such subarray, return 0 instead.

// Input:

// An integer n (1 ≤ n ≤ 10^5): the number of elements in the array.
// An integer target (1 ≤ target ≤ 10^9): the target sum.
// An array nums of length n where 1 ≤ nums[i] ≤ 10^5.
// Output:
// Return the minimal length of a contiguous subarray of which the sum is at least target. If no such subarray exists, return 0.

// Example:

// Input:
// nums = [2, 3, 1, 2, 4, 3]
// target = 7

// Output:
// 2 (The subarray [4, 3] has the minimal length of 2 with a sum of 7.)

package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Create an auxiliary array to store cumulative sums
	cumulativeSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		cumulativeSum[i] = cumulativeSum[i-1] + nums[i-1]
	}

	minLength := math.MaxInt32 // Initialize to a large value
	// Use two pointers to find the minimum length
	for start := 0; start < n; start++ {
		for end := start; end < n; end++ {
			// Calculate the sum of the subarray nums[start:end]
			sum := cumulativeSum[end+1] - cumulativeSum[start]
			if sum >= target {
				// Update the minimum length if the current sum meets the target
				minLength = min(minLength, end-start+1)
			}
		}
	}

	if minLength == math.MaxInt32 {
		return 0 // Return 0 if no valid subarray was found
	}
	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7

	fmt.Println(minSubArrayLen(target, nums))
}
