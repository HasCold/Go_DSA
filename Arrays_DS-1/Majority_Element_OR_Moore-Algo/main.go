package main

import "fmt"

func bruteForceApp(nums []int) int { // Time Complexity => O(n^2)
	n := len(nums)

	for _, tar := range nums {
		count := 0
		for _, num := range nums {
			if tar == num {
				count++
			}
		}

		if count > n/2 {
			return tar
		}
	}

	return -1
}

// Bubble Sort Algorithm
func sorting(nums *[]int) *[]int { // Time Complexity :=  O(n)
	end := len(*nums) - 1

	for i := 0; i < end; i++ {
		if (*nums)[i] > (*nums)[end] {
			temp := (*nums)[i]
			(*nums)[i] = (*nums)[end]
			(*nums)[end] = temp
			end--
		}
	}

	return nums
}

func sortingApproach(nums *[]int) []int { // Time Complexity => O(nlogn)
	sortedArr := sorting(nums)
	fmt.Println("--- Sorted Array ---", *sortedArr)
	n := len(*sortedArr)
	majorElem := []int{}

	frequency := 1
	for i := 0; i < len(*sortedArr)-1; i++ {

		if (*sortedArr)[i] == (*sortedArr)[i+1] {
			frequency++
		} else {
			frequency = 1
		}

		if frequency > n/2 {
			majorElem = append(majorElem, (*sortedArr)[i])
			return majorElem
		}

	}

	return majorElem
}

func mooreAlgo(nums *[]int) int { // Time Complexity O(n) bc there is no any sorting or additional nested loop
	n := len(*nums)
	frequency := 0
	ans := 0

	for i := 0; i < n; i++ {
		if frequency == 0 {
			ans = (*nums)[i]
		}
		if ans == (*nums)[i] {
			frequency++
		} else {
			frequency--
		}
	}

	return ans
}

func main() {
	nums := []int{1, 2, 2, 1, 1}

	fmt.Println("------------- Brute Force FOR Majority Element --------------------")
	fmt.Println("The Brute-Force Approach :- ", bruteForceApp(nums))

	fmt.Println("------------- Sorting Approach FOR Majority Element --------------------")
	nums1 := []int{2, 1, 2, 2, 1}
	fmt.Println("The Sorting Approach :- ", sortingApproach(&nums1))

	fmt.Println("------------- Moore's Voting Algorithm FOR Majority Element --------------------")
	numArr := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("The Moore's Voting Approach :- ", mooreAlgo(&numArr))

	// Moore's Voting Algo Intuition
	// same element =>> Frequency ++
	// different element =>> Frequency --
	// Means if we have a one major element into the array so we can subtract rest of the different elements from them so at the end we will get that majority element from the array

}
