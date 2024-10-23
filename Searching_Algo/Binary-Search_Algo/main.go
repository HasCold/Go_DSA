// Binary Search Alogrithm is the more optimized approach as compared to the Linear Search Algorithm O(n)
// 1. Because in Binary Search Algorithm the Time Complexity is O(logn) which is more efficient as compared to the O(n)
// 2. Binary Search Algorithm can only be apply on the sorted array either in ascending or descending
// e.g. :- [1,2,3,5,10] or [10,9,8,4,2]  -->> Sorted Arrays

package main

import (
	"fmt"
)

func binarySearch(arr []int, tar int) (int, int) {
	// Binary Search Algo can only be apply on the sorted array either ascending or descending
	// mid point := (start point + end point) / 2

	st := 0
	end := len(arr) - 1

	for st <= end {
		// mid := (st + end) / 2  // Prevent the overflow of integer from the large index value
		// To Optimize the code we use this formula
		mid := st + (end-st)/2

		if arr[mid] == tar {
			return arr[mid], mid
		} else if arr[mid] > tar {
			end = mid - 1
		} else if arr[mid] < tar {
			st = mid + 1
		}
	}

	return -1, -1
}

func recursiveBinaryApproach(arr []int, tar int, st int, end int) (int, int) {

	if st <= end {
		// To optimize the code we use this formula
		mid := st + (end-st)/2

		if arr[mid] == tar {
			return arr[mid], mid
		} else if arr[mid] > tar {
			return recursiveBinaryApproach(arr, tar, st, mid-1)
		} else if arr[mid] < tar {
			return recursiveBinaryApproach(arr, tar, mid+1, end)
		}

	}

	return -1, -1
}

func main() {

	arr := []int{-1, 0, 3, 4, 5, 9, 12} // Already Sorted Array
	target := 0

	tarNum, i := binarySearch(arr, target)
	fmt.Printf("The target number %v at index : %v \n", tarNum, i)

	st := 0
	end := len(arr) - 1
	arr1 := []int{-20, -3, 0, 3, 8, 9, 15}
	target2 := 9
	num, index := recursiveBinaryApproach(arr1, target2, st, end)
	fmt.Printf("The recursive binary approach %v : %v  ", num, index)

}
