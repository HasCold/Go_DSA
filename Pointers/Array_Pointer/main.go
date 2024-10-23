package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println("The arr :- ", arr)

	var a = 10 // Suppose memory address => 0x100 (Hexa-decimal form)
	var p *int = &a
	fmt.Println("The Before a :- ", a)   // 10
	fmt.Println("The Before p++ :- ", p) // 0xc000102080

	// Go does not support pointer arithmetic like p++.
	// p++
	*p++
	fmt.Println("The After p++ of a :- ", a) // 11
	fmt.Println("The After p++ of a :- ", p) // 0xc000102080

}
