// ------------------------------- Search in Rotated Sorted Array ------------------------------------
// There is an integer array nums sorted in ascending order (with distinct values).

// Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

// Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [3,4,5,6,7,0,1,2], target = 0
// Output: 4

// Approach :-
// We have to pick the "search space" very cautiously in the binary search algo because in this case we have a rotated_sorted_array.
// 1. When we talk about the rotated sorted array there always be a left-sorted-array or either be a right-sorted-array.
// 2. e.g. :- [3,4,5,6,7,0,1,2] ; mid = 6;  so as you can see that our left-search-space is already sorted [3,4,5] rather than right-search-space [7,0,1,2]
// 3. we have left-search-space sorted so we can apply the binary search condition on that.

package main

import "fmt"

func rotatedSortedArray(nums []int, tar int) (int, int) { // Time Complexity O(logn)

	st := 0
	end := len(nums) - 1

	for st <= end {
		// To prevent the overflow of integer by the large size index
		// To optimize the code we use this formula
		mid := st + (end-st)/2

		if nums[mid] == tar {
			return nums[mid], mid
		} else if nums[st] <= nums[mid] { // left-search-space check either sorted or not

			if nums[st] <= tar && tar <= nums[mid] { // checks if target exist between the left sorted search space
				end = mid - 1
			} else { // then checks if target exist between the right search space
				st = mid + 1
			}

		} else if nums[end] >= nums[mid] { // right-search-space check either sorted or not

			if nums[mid] <= tar && tar <= nums[end] { // checks if target exist between the right sorted search space
				st = mid + 1
			} else { // then checks if target exist between the left search space
				end = mid - 1
			}
		}
	}

	return -1, -1
}

func main() {
	nums := []int{3, 4, 5, 6, 7, 0, 1, 2}
	target := 0

	num, i := rotatedSortedArray(nums, target)
	fmt.Printf("Rotated Sorted Array of target %v at index %v", num, i)

}
