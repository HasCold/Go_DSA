// Single Element in a Sorted Array

// You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.

// Return the single element that appears only once.

// Your solution must run in O(log n) time and O(1) space.

// Example 1:

// Input: nums = [1,1,2,3,3,4,4,8,8]
// Output: 2
// Example 2:

// Input: nums = [3,3,7,7,10,11,11]
// Output: 10

// Constraints:
// 1 <= nums.length <= 10^5
// 0 <= nums[i] <= 10^5

// Approach :-
// mid of the array [3,3,7,7,10,11,11] = mid-index = 3 (odd) so our non-duplicate number will be present at right-search-space of the array
// mid of the array [1,1,2,3,3,4,4,8,8] = mid-index = 4 (even) so our non-duplicate number will be present at left-search-space of the array

package main

import "fmt"

func singleNonDuplicate(arr []int) int { // O(logn) Time Complexity and O(1) Space Complexity

	st := 0
	end := len(arr) - 1

	for st <= end {
		mid := st + (end-st)/2

		// ---------------- Handling Edge Cases -----------------
		if len(arr) == 1 {
			return arr[mid]
		}
		if mid == 0 && arr[mid] != arr[mid+1] {
			return arr[mid]
		}
		if mid == len(arr)-1 && arr[len(arr)-1] != arr[len(arr)-2] {
			return arr[len(arr)-1]
		}
		// ---------------- Handling Edge Cases -------------------

		if arr[mid] != arr[mid-1] && arr[mid] != arr[mid+1] {
			return arr[mid]

		} else if mid%2 == 0 { // If our mid index is even then the left-search-space and right-search-space size will also even
			if arr[mid] == arr[mid-1] {
				end = mid - 1
			} else if arr[mid] == arr[mid+1] {
				st = mid + 1
			}

		} else if mid%2 != 0 { // If our mid index is odd then the left-search-space and right-search-space size will also odd
			if arr[mid] == arr[mid-1] {
				st = mid + 1
			} else if arr[mid] == arr[mid+1] {
				end = mid - 1
			}
		}
	}

	return -1
}

func main() {
	nums1 := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	nums2 := []int{3, 3, 7, 7, 10, 11, 11}

	fmt.Println("To find the non-duplicate number in the even-mid array size", singleNonDuplicate(nums1))
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("To find the non-duplicate number in the odd-mid array size", singleNonDuplicate(nums2))
}
