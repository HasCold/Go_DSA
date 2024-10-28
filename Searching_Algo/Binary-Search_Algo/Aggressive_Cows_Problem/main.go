// ------------------------------ Aggresive Cows Problem -----------------------------------

// Assign C rows to N stalls such that a minimum distance between them is largest possible.
// Return the largest minimum distance

// Find the larget minimum distance between each cows. Permutation

package main

import (
	"fmt"
	"slices"
)

// Possible ranges to maintain the min distance between cows
// 1,2,3,4,5,6,7,8,9
// min = 1			max= largest in arr - smallest in arr = 8
// largest min distance is going to be in between 1 to 8
func largestMinDistance(arr []int, cows int) int {

	if cows > len(arr) {
		return -1
	}

	// minDistance := 1
	// Sort the slice
	slices.Sort(arr) // O(nlogN)

	st := 0
	end := arr[len(arr)-1] - arr[0] // last ele - first ele = 8
	ans := -1

	for st <= end { // O(log(Range) * N)

		// To prevent the overflow of integer by the large size indexes
		// We optimize the code by using this formula
		mid := st + (end-st)/2
		fmt.Println(isPossible(mid, arr, cows))

		if isPossible(mid, arr, cows) { // if condition false then we find out in the right-seaarch-space
			ans = mid
			st = mid + 1
		} else { // if condition becomes true then we find out in the left-seaarch-space
			end = mid - 1
		}
	}

	return ans
}

// 1,2,3,4,5,6,7,8,9
// {1, 2, 4, 8, 9}
func isPossible(minAllowedDistance int, arr []int, cows int) bool { // O(N)

	cow := 1
	lastStall := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i]-lastStall >= minAllowedDistance {
			cow++
			lastStall = arr[i]
		}

		if cow == cows {
			return true
		}
	}

	return false
}

func main() {
	arr := []int{1, 2, 8, 4, 9} // No. of position w.r.t stalls
	// N := len(arr)               // No. of stalls
	C := 3 // No. of cows

	fmt.Println("The largest minimum distance between the cows is :- ", largestMinDistance(arr, C))
}
