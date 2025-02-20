// Sort an array of 0s, 1s and 2s
// DNF algorithm

package main

import "fmt"

// Insertion Sort
func bruteForceApproach(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		prev := i - 1

		for prev >= 0 && arr[prev] > cur {
			arr[prev+1] = arr[prev]
			prev--
		}
		arr[prev+1] = cur
	}

	return arr
}

// Optimized Approach
func optimizeApproach(arr []int) []int { // O(2n)  becomes equal to O(n) Time Complexity

	countOfZeros := 0
	countOfOnes := 0
	countOfTwos := 0

	for i := 0; i < len(arr); i++ { // O(n)
		if arr[i] == 0 {
			countOfZeros++
		} else if arr[i] == 1 {
			countOfOnes++
		} else {
			countOfTwos++
		}
	}

	for i := 0; i < len(arr); i++ { // O(n)
		if countOfZeros > 0 {
			arr[i] = 0
			countOfZeros--
		} else if countOfZeros == 0 && countOfOnes > 0 {
			arr[i] = 1
			countOfOnes--
		} else if countOfOnes == 0 && countOfTwos > 0 {
			arr[i] = 2
			countOfTwos--
		}
	}

	return arr
}

// More Optimal Approach  -->> DNF Algo (Dutch National Flag Algorithm)

// Approach :-
// 0 	  l-1 	l	  m-1   m				   h	h+1	  n-1
// <- 0000 ->  <- 1111 ->  <- unsorted elements ->  <- 222 ->
// We are using three poniter approach : low, mid and high

func dnfAlgo(arr []int) []int { // O(n)  Time Complexity and O(1) Space Complexity

	low := 0
	mid := 0
	high := len(arr) - 1

	//  l
	// [2,0,2,1,1,0]
	//  m		  h

	for mid <= high {
		if arr[mid] == 0 { // Condition for 0's
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 { // Condition for 1's
			mid++
		} else { // Condition for 2's
			arr[high], arr[mid] = arr[mid], arr[high]
			high--
		}
	}

	return arr
}

func main() {
	arr := []int{2, 0, 2, 1, 1, 0, 1, 2, 0, 0}

	arr2 := []int{2, 0, 2, 1, 1, 0, 1, 2, 0, 0}

	fmt.Println("The brute-force approach for 0s, 1s and 2s :- ", bruteForceApproach(arr))

	fmt.Println("The optimized approach for 0s, 1s and 2s :- ", optimizeApproach(arr2))

	arr3 := []int{2, 0, 2, 1, 1, 0, 1, 2, 0, 0}
	fmt.Println("The optimal approach for 0s, 1s and 2s :- ", dnfAlgo(arr3))
}
