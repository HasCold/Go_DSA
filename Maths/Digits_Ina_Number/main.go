package main

import "fmt"

// The 3568 is a Base-10 decimal-number-system could be divided by 10 for %

// n 		=   	 3586 % 10 = 6 (Remainder)
// quotient	= 		 358 % 10 = 8 (Remainder)
// quotient	= 		 35 % 10 = 5 (Remainder)
// quotient	= 		 3 % 10 = 3 (Remainder)
// 			=         0

func printDigits(n int) { // Time Complexity := O(log n)  becuase after decreasing our n value by 10 so it will log base 10
	count := 0
	var sum int

	for n != 0 {
		digit := n % 10 // Remainder //
		fmt.Println(digit)
		n = n / 10 // Quotient  // log base 10

		sum += digit
		count++
	}
	fmt.Println("The total digit count :- ", count)
	fmt.Println("The Sum of digit :- ", sum)
}

func main() {
	var n = 3586
	printDigits(n)
}
