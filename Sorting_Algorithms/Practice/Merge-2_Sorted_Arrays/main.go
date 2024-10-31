package main

import "fmt"

// nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3

func conquer(arr []int, st, mid, end int) {
	// Sort the arrays and then merged two sorted arrays

	// Creating a Slices
	left := arr[st : mid+1]     // Covers from 0 to mid
	right := arr[mid+1 : end+1] // Covers from mid + 1 index to end index

	mergedArr := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			mergedArr = append(mergedArr, left[i])
			i++
		} else {
			mergedArr = append(mergedArr, right[j])
			j++
		}
	}

	for i < len(left) {
		mergedArr = append(mergedArr, left[i])
		i++
	}

	for j < len(right) {
		mergedArr = append(mergedArr, right[j])
		j++
	}

	copy(arr[st:end+1], mergedArr)
}

func divide(arr1 []int, st, end int) {

	if st >= end {
		return
	}

	mid := st + (end-st)/2

	divide(arr1, st, mid)
	divide(arr1, mid+1, end)
	conquer(arr1, st, mid, end)

}

func merge2SortArr(arr1 []int, arr2 []int, m, n int) []int {

	// Divide & Conquer
	sort := directConquer(arr1, arr2, m, n)
	return sort
}

func directConquer(arr1, arr2 []int, m, n int) []int {

	mergedSortArr := make([]int, 0, m+n)
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			mergedSortArr = append(mergedSortArr, arr1[i])
			i++
		} else {
			mergedSortArr = append(mergedSortArr, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		mergedSortArr = append(mergedSortArr, arr1[i])
		i++
	}

	for j < len(arr2) {
		mergedSortArr = append(mergedSortArr, arr2[j])
		j++
	}

	// copy(arr1[0:m+n], mergedSortArr)
	return mergedSortArr
}

func main() {
	// Divide & Conquer Approach
	nums1 := []int{2, 2, 1, 0, 0, 0}
	m := 3 // the non-zero element in num1
	nums2 := []int{5, 6, 3}
	n := len(nums2)
	mergedArr := make([]int, 0, m+n)
	mergedArr = append(nums1[:m], nums2...) // means 0 to 2 index

	st1 := 0
	end1 := len(mergedArr) - 1
	divide(mergedArr, st1, end1)
	fmt.Println("The merged sorted arrays is :- ", mergedArr)

	//-----------------------------------------------------------------------------------------------------
	num3 := []int{1, 2, 3, 0, 0, 0}
	p := 3           // non-zero element in num3
	arr1 := num3[:p] // 0 to 2 index
	arr2 := []int{2, 5, 6}
	q := len(arr2)

	merged2SortedArr := merge2SortArr(arr1, arr2, p, q)

	fmt.Println("Merged Two Sorted Arrays :- ", merged2SortedArr)
}
