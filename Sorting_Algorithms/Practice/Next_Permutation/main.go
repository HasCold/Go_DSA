package main

import (
	"fmt"
)

// Permutation => Means the number of arrangement or no. of order
// For example, for arr = [1,2,3], the following are all the permutations of arr:
// [1,2,3],
// [1,3,2],
// [2,1,3],
// [2,3,1],
// [3,1,2],
// [3,2,1].

// Return lexicographically next
// Means lexicographically :- Order of arrangement from smaller to bigger.

// Brute Force Approach
// 1. Find all permutations
// 2. Find lexicographically next permutation
// 3. Recursion
// 4. n = 3  means 6 permutations / 3 factorial (3! = 3*2*1 = 6)
// 5. O(n! x n)  => n-factorial

// Optimal Approach :-
// 1. O(1) Space Complexity
// 2. O(n) Time Complexity

// 1,2,3,5,4  -->> 1,2,4,3,5   // 3 < 5

// 1,2,3,6,5,4  -->> 1,2,4,3,5,6   // 3 < 6

// 1,2,5,4,3  -->> 1,3,2,4,5   // 2 < 5

// Find the pivot element => A[i] < A[i+1] from the backward
// Find the right most(backward) element that will greater than pivot
// Swap right most element , pivot element
// Reverse Elements from pivot+1 to n-1

func optimalApproach(arr []int) { // Time Complexity => O(3n) => O(n), Space Complexity => O(1)

	var pivot int
	var pvtidx int
	isFound := false

	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] { // A[i] < A[i+1]
			pivot = arr[i]
			pvtidx = i
			isFound = true
			break
		}
	}
	if !isFound {
		reverse(arr, 0, len(arr)-1)
		return
	}

	for i := len(arr) - 1; i > pvtidx; i-- {
		if arr[i] > pivot {
			arr[i], arr[pvtidx] = pivot, arr[i]
			break
		}
	}

	reverse(arr, pvtidx+1, len(arr)-1)
}

func reverse(arr []int, stIdx, endIdx int) {

	for stIdx < endIdx {
		arr[stIdx], arr[endIdx] = arr[endIdx], arr[stIdx]
		stIdx++
		endIdx--
	}
}

func main() {
	A := []int{3, 2, 1}
	B := []int{1, 2, 3, 6, 5, 4}

	optimalApproach(A)
	fmt.Printf("The lexicographically next permutation is :- %v", A)

	fmt.Println()
	optimalApproach(B)
	fmt.Printf("The lexicographically next permutation is :- %v", B)

}
