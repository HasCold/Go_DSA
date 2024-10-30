package main

import "fmt"

// Merge Sort :-
// Approach works on Divide and Conquer Rule

func partitioning(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	// {6, 3, 9, 5, 2, 8}
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[i], arr[high] = arr[high], arr[i]

	return i
}

func quickSort(arr []int, low, high int) []int {
	// Partitioning basis on pivot point
	// For quick sort implmentation we select the smallest element as a pivot point which is the last element index

	// 1. First select the pivot point by assume the last index is our pivot index which is the smallest element considering
	// 2. Then place the pivot element at their right order according to the sorted array

	if low < high {
		pvtIdx := partitioning(arr, low, high)

		quickSort(arr, low, pvtIdx-1)  // To sort the smaller elements which is of right-side of array
		quickSort(arr, pvtIdx+1, high) // To sort the larger elements
	}

	return arr
}

func conquer(arr []int, st, mid, end int) {

	//  Create a slice
	left := arr[st : mid+1]     // 0 - mid
	right := arr[mid+1 : end+1] // mid+1 - end

	merged := make([]int, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	for i < len(left) { // Print all the remaining elements of left array to merged array
		merged = append(merged, left[i])
		i++
	}

	for j < len(right) { // Print all the remaining elements of right array to merged array
		merged = append(merged, right[j])
		j++
	}

	copy(arr[st:end+1], merged) // array length => st(0) to end(n)
}

// Divide and Conquer Approach For Merge Sort
func divide(arr []int, st, end int) {

	if st >= end {
		return
	}

	// To prevent the overflow of integer by the large size indexes so we use this formula to optimize the code
	mid := st + (end-st)/2

	divide(arr, st, mid) // O(logn) Time Complexity
	divide(arr, mid+1, end)
	conquer(arr, st, mid, end)

}

func main() {

	arr := []int{6, 3, 9, 5, 2, 8}
	n := len(arr) - 1

	sort := quickSort(arr, 0, n)
	fmt.Println("The sorted array from quick sort algo :- ", sort)

	arr2 := []int{6, 3, 9, 5, 2, 8}
	st := 0
	end := len(arr2) - 1
	divide(arr2, st, end)
	fmt.Println("The sorted array from merge sort algo :- ", arr2) // Overall Time Complexty is O(nlogn)

}
