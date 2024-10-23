// Null Pointer :- A pointer that doesn't point to any location.

// We can pass the argument in the function by two ways :-
// 1. Pass by value (Copy)
// 2. Pass by reference / alias (actual memory address location shared of that data)
// 3. Pass by pointers (Means creating a new var of pointer in which they stored a same memory address location)

package main

import "fmt"

func main() {
	// var ptr *int   // If we don't assign the nil / null so we get the new garbage value every time
	var ptr *int = nil

	fmt.Println("The NULL Pointer :- ", &ptr) //  0xc00004a028
	// fmt.Println("The NULL Pointer :- ", *ptr) // Error := Segmentation Fault / Access that space that is not available

	var a int = 10
	var p *int = &a
	var q **int = &p

	fmt.Println(*p)  // 10
	fmt.Println(**q) // 10
	fmt.Println(p)   // 0xc000092010
	fmt.Println(*q)  // 0xc000092010
}
