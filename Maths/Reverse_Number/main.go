package main

import (
	"fmt"
	"math"
)

// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:

// Input: x = 123
// Output: 321

func reverseNum(num int) int {

	var reversedNum = 0

	for num != 0 {
		digit := num % 10 // remainder digit

		// Check that our reversed Number is in the range of 32-bit opearting ; not out of the bound
		if (reversedNum > math.MaxInt32/10) || (reversedNum < math.MinInt32/10) {
			return 0
		}
		reversedNum = (reversedNum * 10) + digit // 70 + 3 = 73

		num = num / 10 // quotient

	}

	return reversedNum
}

func main() {
	num := 120 // Output -->> 21
	num1 := 1534236469
	num2 := -4537 // Output -->> 7354-

	fmt.Println(reverseNum(num)) // 7354
	fmt.Println(reverseNum(num1))
	fmt.Println(reverseNum(num2))
}
