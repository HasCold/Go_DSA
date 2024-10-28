// You are given a list of n houses located on a straight line at distinct integer coordinates. Your goal is to place k routers in such a way that the maximum distance from any house to its nearest router is minimized.

// Input:
// An integer n (1 ≤ n ≤ 100,000): the number of houses.
// An integer k (1 ≤ k ≤ n): the number of routers to be placed.
// An array houses of length n where houses[i] (0 ≤ houses[i] ≤ 1,000,000,000) is the coordinate of the i-th house.

// Output:
// Return the minimized maximum distance (an integer) from any house to its nearest router.

// Example:

// Input:
// houses = [1, 2, 8, 4, 9]
// k = 3

// Output:
// 3

package main

import (
	"fmt"
	"slices"
)

func isPossible(maxPossibleDistance int, arr []int, k int) bool {

	house := arr[0]
	lastPos := 1
	// [1,2,4,8,9]
	for i := 1; i < len(arr); i++ {
		if arr[i]-house >= maxPossibleDistance {
			house = arr[i]
			lastPos++
		}
	}
	// If we can place at least k routers, return true
	return lastPos >= k
}

func minimizedDistance(arr []int, k int) int {

	if k > len(arr) {
		return -1
	}

	slices.Sort(arr)

	// minDistance = 1
	// maxDistance = 9-1 = 8

	st := arr[0]
	end := arr[len(arr)-1] - arr[0] // largest distance - smallest distance = 9-1 = 8
	ans := -1
	// [1,2,4,8,9]  -->> Now Apply Binary Search Algorithm
	for st <= end {
		// To prevent the overflow of integer by the large size indexes we use this formula to optimiize the code
		mid := st + (end-st)/2

		if isPossible(mid, arr, k) {
			ans = mid
			st = mid + 1 // Try for a larger distance
		} else {
			end = mid - 1 // Try for a smaller distance
		}

	}

	return ans
}

func main() {
	houses := []int{1, 2, 8, 4, 9}
	k := 3 // No. of routers

	fmt.Println("The minimized maximum distance from any house to its nearest integer :- ", minimizedDistance(houses, k))
}
