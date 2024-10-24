// Peak Index in a Mountain Array

// You are given an integer mountain array arr of length n where the values increase to a peak element and then decrease.
// Return the index of the peak element.

// Your task is to solve it in O(log(n)) time complexity.

// Example 3:
// Input: arr = [0,10,5,2]

// Output: 1

// So it means we have to solve them from the binary search algorithm.

package main

import (
	"errors"
	"fmt"
)

func peakIndexInMountainArray(arr []int) (int, int, error) {
	// First we already know the peak value doesn't exit at the mountain array start or end index.
	// so we initialize our start and end point from 1 and n-1

	st := 1
	end := len(arr) - 2

	for st <= end {
		// To prevent the overflow of integer by the large size index
		// To optimize the code we use this formula
		mid := st + (end-st)/2

		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] { // check whether the coming mid point element is peak index or not
			return arr[mid], mid, nil

		} else if arr[mid-1] < arr[mid] { // just only check in the left-search space the mid value are present in increasing side or not then we must find the peak value in the right-search-space
			st = mid + 1

		} else if arr[mid+1] < arr[mid] { // just only check in the right-search space the mid value are present in decreasing side or not then we must find the peak value in the left-search-space
			end = mid - 1
		}

	}

	return -1, -1, errors.New("peak index doesn't present in the mountain array.")
}

func main() {
	arr := []int{0, 3, 8, 9, 5, 2}

	ele, i, err := peakIndexInMountainArray(arr)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The peak index from the mountain array of element : %v at index %v", ele, i)
	}
}
