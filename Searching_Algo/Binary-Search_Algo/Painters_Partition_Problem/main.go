// Question :-

// Given N boards of length of each given in the form of array and M painters, such that each painter takes 1 unit of time to paint 1 unit of board

// The task is to find the minimum time to paint all boards under the constraints that any painter will only paint continuous sections of boards.

// Approach :-
// 1. arr = [40, 30, 10, 20] ; M = 2 painter's
// 2. Permutations :-
//  P1 = 40 ; P2 = 60  // The minimum time to paint all the no. of blocks in a unit time
//  P1 = 70 ; P2 = 30
//  P1 = 80 ; P2 = 20

package main

import "fmt"

// [10, 20, 30, 40, 50, 60, 70, 80, 90, 100] -->> Apply the Binary Search Algo on sorted array for min time taken to paint the block
// Min = 10 and Max no. of boards = 100 so the minimum time can only be lie between these two ranges
func partitionProblem(arr []int, painter int) int { // Time Complexity = O(logn) + O(n) => O(nlogN)

	// Edge case
	if painter > len(arr) {
		return -1
	}
	// Edge case

	st := 0
	end := sum(arr)
	var ans int

	for st <= end {
		mid := st + (end-st)/2

		if isValid(arr, painter, mid) { // if isValid = true then we move the operation in the left-search-space
			ans = mid
			end = mid - 1
		} else { // if isValid = false then we move the operation in the rigth-search-space
			st = mid + 1
		}

	}

	return ans
}

func isValid(arr []int, painter int, maxAllowableTime int) bool {

	totalPaints := 0
	worker := 1
	// {40, 30, 10, 20}
	for i := 0; i < len(arr); i++ {
		if totalPaints+arr[i] <= maxAllowableTime {
			totalPaints += arr[i]
		} else {
			worker++
			totalPaints = arr[i]
		}

	}
	if worker <= painter {
		return true
	} else {
		return false
	}

}

func sum(arr []int) int {
	ans := 0
	for _, val := range arr {
		ans += val
	}

	return ans
}

func main() {
	arr := []int{40, 30, 10, 20} // No. of wooden blocks
	// N := len(arr)

	M := 2 // No. of painters who has worked that they paints the wooden blocks

	fmt.Println("The minimum time taken by painter's to paint all the blocks :- ", partitionProblem(arr, M))
}
