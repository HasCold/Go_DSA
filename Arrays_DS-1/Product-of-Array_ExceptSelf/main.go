// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// Constraints :-
// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// Approach :=
// 1. nums = [1,2,3,4] , ans = [24,12,8,6]
// we can do store all the product of element into a var = 24, then we can divide them with nums[i] so we got this ans = [24,12,8,6]

// Array or Slice are passed by the reference in the function

package main

import "fmt"

// prefix  -->> First word
// suffix -->> Last word
// (Prefix) leftProd <-- index i --> (Suffix) RightProd

func productExceptSelf(nums []int) []int { // Time Complexity O(n)
	// Array or Slice are passed by the reference in the function rather than passed by value(copy)
	// This change won't affect the original slice
	ans := make([]int, len(nums))
	// copy(ans, nums)

	// prefix := make([]int, len(nums))
	// prefix[0] = 1
	// prefix := 1 // More Optimal Approach

	// suffix := make([]int, len(nums))
	// suffix[len(nums)-1] := 1
	suffix := 1 // More Optimal Approach
	ans[0] = 1

	// O(3n)  ==>> which will amount to O(n) Time Complexity

	// prefix ; first index of prefix is already calculated
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	// suffix ; last index is already calculated
	for i := len(nums) - 2; i >= 0; i-- {
		suffix = suffix * nums[i+1]
		ans[i] *= suffix
	}

	return ans
}

func bruteForceApproach(nums []int) []int { // Time Complexity O(n^2) and Space Complexity O(1)

	// Array or Slice are passed by the reference in the function rather than passed by value(copy)
	// This change won't affect the original slice
	ans := make([]int, len(nums))
	copy(ans, nums)

	for i := 0; i < len(nums); i++ {
		var prod = 1
		for j := 0; j < len(nums); j++ {
			if nums[i] != nums[j] {
				prod *= nums[j]
			}
		}
		ans[i] = prod
	}

	return ans
}

func main() {

	nums := []int{1, 2, 3, 4}
	// productExceptSelf(nums)

	var i = bruteForceApproach(nums)
	fmt.Println("The brute force approach is :=", i)

	fmt.Println("------------- The Optimize Approach --------------")
	fmt.Println(productExceptSelf(nums))
}
