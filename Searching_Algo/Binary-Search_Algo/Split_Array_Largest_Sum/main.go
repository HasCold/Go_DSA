// Given an array nums which consists of n integers, split the array into m non-empty continuous subarrays.

// Write a function to minimize the largest sum among these m subarrays.

// Example:

// plaintext
// Copy code
// Input: nums = [7, 2, 5, 10, 8], m = 2
// Output: 18
// Explanation:
// There are several ways to split the array into two subarrays.
// - [7, 2, 5] and [10, 8], with sums 14 and 18. The largest sum is 18.
// - [7, 2, 5, 10] and [8], with sums 24 and 8. The largest sum is 24.

// The minimum largest sum among these splits is 18.
// Constraints:

// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 10^6
// 1 <= m <= min(50, nums.length)

package main

import "fmt"

func maxInArray(arr []int) int {
	maxVal := arr[0]

	for _, num := range arr {
		if num > maxVal {
			maxVal = num
		}
	}

	return maxVal
}

func sumArray(arr []int) int {
	currSum := 0

	for _, num := range arr {
		currSum += num
	}

	return currSum
}

func isValid(maxSum int, arr []int, split int) bool {

	currSum := 0
	subArray := 1

	for i := 0; i < len(arr); i++ {
		if currSum+arr[i] > maxSum {
			subArray++
			currSum = arr[i]

			if subArray > split {
				return false
			}

		} else {
			currSum += arr[i]
		}
	}

	return true
}

// min(10) <------------------> max(32)  // Binary Search Algo can be now apply on these sorted ranges
func splitArray(arr []int, split int) int {
	st, end := maxInArray(arr), sumArray(arr)
	ans := -1

	for st <= end {
		// To prevent the overflow of integer by the large size indexes we use this formula
		mid := st + (end-st)/2

		if isValid(mid, arr, split) { // if true then we find in the left-search-space
			ans = mid
			end = mid - 1
		} else { // if false then we find in the right-search-space
			st = mid + 1
		}

	}

	return ans
}

func main() {
	nums := []int{7, 2, 5, 10, 8} // maximum subarray
	m := 2                        // split the array into two

	// Write a function to minimize the largest sum among these m subarrays.
	fmt.Println("The minimize largest sum :- ", splitArray(nums, m))
}
