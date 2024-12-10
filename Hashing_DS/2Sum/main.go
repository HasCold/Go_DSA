package main

import (
	"fmt"
	"slices"
)

// We already know that binary search algorithm will be implement on the sorted array

// 1. Brute Approach
// O(n^2)

//2. Better Approach
//	Sort   +   Operation
// O(n) + O(nlogn)  ==>> O(nlogn)

func TwoSum(arr []int, tar int) []int {

	// Sort out the array
	slices.Sort(arr)

	st := 0
	end := len(arr) - 1
	var sum int
	var ans []int

	// 2 pointer approach
	for st < end {
		sum = arr[st] + arr[end]
		if sum == tar {
			ans = append(ans, st, end)
			break
		}

		if sum > tar {
			end--
		} else if sum < tar {
			st++
		}
	}

	return ans
}

// Optimized Approach by Hashing DS used
func TwoSumOptimized(arr []int, target int) []int { // O(n)
	pair := map[int]int{} // stores the key-value pair
	ans := []int{}

	// sum = firstIdx + secondIdx == target

	for i := 0; i < len(arr); i++ {
		first := arr[i]
		second := target - first

		_, exist := pair[second]
		if exist {
			ans = append(ans, i, pair[second])
			break
		}

		// Store the key-value pair , implemented a Hash Table , Providing average 0(1) time complexity
		pair[first] = i
	}

	return ans
}

func main() {
	arr := []int{5, 2, 11, 7, 15}
	target := 9

	fmt.Println("The better approach of two sum is := ", TwoSum(arr, target))

	fmt.Println("The optimized approach of two sum is := ", TwoSumOptimized(arr, target))
}
