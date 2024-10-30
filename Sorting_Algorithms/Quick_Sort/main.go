// ### 2. **Efficient Sorting Algorithms**
//    - **Merge Sort**: A divide-and-conquer algorithm that splits the array, sorts, and merges. Time complexity is O(n log n).
//    - **Quick Sort**: Also a divide-and-conquer algorithm, but uses a pivot to partition the array. Generally faster than merge sort but has O(n²) worst-case time complexity (usually avoided by choosing good pivots).
//    - **Heap Sort**: Uses a binary heap data structure to sort elements. Has O(n log n) time complexity and is in-place, but not stable.

// Merge , Quick and Heap Sort maninly works on your O(nlogn) Time Complexity

package main

import "fmt"

// Quick Sort Intution -->>
// - Partitioning Logic
// - Quick sort implemented around the pivot point
// - There are several ways to choose the pivot point but we are selecting the pivot point to our last index element.
// - Then place the pivot according to the right place

// e.g. :- [6,3,9,5,2,8]
// pivot is the last index ele -> 8 so place the 8 in the right order

func partition(arr []int, low, high int) int { // will return at this final recursive sort [2,3,5,6,8,9]

	pivot := arr[high] // consider pivot array of last element
	i := low - 1       // index tracking

	for j := low; j < high; j++ {
		if arr[j] < pivot { // This logic will sort the smaller element move left-side by the pivot(8)
			// Swap
			i++
			arr[i], arr[j] = arr[j], arr[i]

		}
	}
	// Place the pivot(8) element at its correct position
	i++
	arr[i], arr[high] = arr[high], arr[i]

	return i
}

func quickSort(arr []int, low int, high int) []int {

	// 1. First select the pivot point by assume the last index is our pivot index which is the smallest element considering
	// 2. Then place the pivot element at their right order according to the sorted array

	if low < high {
		var pidx = partition(arr, low, high) // Partition function give us the right position of pivot index

		quickSort(arr, low, pidx-1)  // means if index at which 8 is pivot so before that index all the elements are unsorted.
		quickSort(arr, pidx+1, high) // means if index at which 8 is pivot so after that index all the elements are unsorted.

	}

	return arr
}

func main() {

	arr := []int{6, 3, 9, 5, 2, 8}
	lastEleInd := len(arr) - 1 // Last Element taken to be as a pivot so that quicksort will implemneted around that pivot

	sortArr := quickSort(arr, 0, lastEleInd)
	fmt.Println("The sorted array by using quick sort :- ", sortArr)

}

// Time Complexity of Quick Sort :-
// Average Case := O(nlogn)
// Worst := O(n^2)
// Important :-
// worst case occurs when pivot is always the smallest or the largest element.

// Most of the time Quick Sort will be preferred as compared to Merge Sort bc of O(1) space complexity in quick sort. And rarely we can see the worst time complexity in quick sort
