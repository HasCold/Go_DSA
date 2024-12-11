package main

import "fmt"

// T.C := O(n)
// S.C := O(n)  because of the use of Hash Table

func findDuplicate(arr []int) int {

	find := map[int]int{} // key, value pair
	var a = -1

	for i := 0; i < len(arr); i++ {
		num := arr[i]
		_, exist := find[num]

		if exist {
			a = num
			return a
		}
		find[num] = i
	}

	return a

}

// Slow-Fast Pointer Approach
// element :-  [3, 1, 3, 4, 2] ==>> Next Node
// index :- 	0, 1, 2, 3, 4  ==>> Node

// Imagine array as a linkedlist

// Practically :- 0, 3, 4, 2, 3
// T.C. := O(n)
// S.C. := O(1)

// 1. Slow Pointer => +1
// 2. Fast Pointer => +2
// Slow == Fast

func slowFastApproach(arr []int) int {

	slow := arr[0]
	fast := arr[0]

	for {
		slow = arr[slow]      // Update with +1
		fast = arr[arr[fast]] // Update with +2

		if slow == fast {
			break
		}
	}

	slow = arr[0]
	for slow != fast {
		slow = arr[slow] // Update with +1
		fast = arr[fast] // Update with +1
	}

	return slow
}

func main() {
	arr := []int{3, 1, 3, 4, 2}

	fmt.Println(findDuplicate(arr))

	fmt.Println("The slow fast pointer approach :- ", slowFastApproach(arr))
}
