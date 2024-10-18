package main

import "fmt"

func pairSum(nums []int, target int) []int { // Time Complexity => O(n^2)
	ans := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans = append(ans, i, j)
			}
		}
	}

	return ans
}

func twoPointerApproch(nums []int, target int) []int { // Time Complexity => O(n)
	i := 0
	j := len(nums) - 1
	ans := []int{}

	for i < j {
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			ans = append(ans, i, j)
			break
		}
	}

	return ans
}

func main() {
	// This is already an ascending sorted array
	var nums = []int{2, 7, 11, 15}
	target := 26

	fmt.Println("--------------- Brute-Force Approach For Pair Sum -------------------")
	indexSlice := pairSum(nums, target)
	fmt.Println(indexSlice)

	fmt.Println("--------------- Two Pointer Approach For Pair Sum -------------------")
	// PairSum(PS)
	// PS > tar  	i--
	// PS < tar 	i++
	// PS == tar	ans => (i,j)
	numIndex := twoPointerApproch(nums, target)
	fmt.Println(numIndex)

}
