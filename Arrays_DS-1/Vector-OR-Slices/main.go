package main

import "fmt"

// Vector or SLices :- In golang we have a slices which is a powerful, flexible and light-weight data-structure. Slices are dynamic in length.

// Slices can be passed by value or passed by reference.

func push(slice []int, value int) []int {
	return append(slice, value)
}

func pop(slice []int) ([]int, int) {
	if len(slice) == 0 {
		return slice, -1
	}

	popValue := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	// [0:4]  means take a slice from 0 to 3 index, []
	// [:]  slices containing all elements
	// [:len(slice)-1]  slice containing element from 0 index to len(slice)-2

	return slice, popValue
}

func singleNumber(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}

	return ans
}

func main() {
	vec := []int{0, 1, 2, 3, 4, 5}

	fmt.Println("The size of vector :-", len(vec))

	value := 23
	pushedVec := push(vec, value)
	fmt.Println("The pushed value of vector :-", pushedVec)

	slice, popValue := pop(vec)
	fmt.Printf("The poppped value : (%v) , their data-type : (%T) and popped slice : (%v) \n", popValue, popValue, slice)

	// Single Number Problem
	var nums = []int{4, 1, 2, 1, 2}

	// We can solve this problem by Bitwise XOR Operator ; solution with a linear runtime complexity (means one loop)
	// same bit = 0
	// diff bit = 1
	// 0 ^ 0 = 0
	// 0 ^ 1 = 1
	// 1 ^ 0 = 1
	// 1 ^ 1 = 0
	uniqueNum := singleNumber(nums)
	fmt.Println("The unique number :- ", uniqueNum)
}

// Constraints will tell about the limit

// ------------------------- Static VS Dynamic Allocation -------------------------

// Code will run at the two stages :=
// 1. Compile time
// 2. Run time / Execution time

// Static Allocation will allot on the compile time ; var num = [5]int
// Dynamic Allocation will allot in memory on the run time ; num := []int

//  There are two types of memory
// 1. Stack
// 2. Heap

// Static Allocation will done in the stack memory e.g. Arrays
// Dynamic Allocation will done in the heap memory e.g. Slices
