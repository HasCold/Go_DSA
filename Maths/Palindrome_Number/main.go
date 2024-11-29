package main

import (
	"fmt"
	"math"
)

// Vald Palindrome String -->> Means when we read the string from forward or from backward it still remains same.
// eg :- race car

// Palindrome Number :- 353 -->> 353  Reverse of a number

func reversed(x int) int {
	// 4537
	reversedNum := 0

	for x != 0 {
		digit := x % 10 // remainder

		// Check that reversedNum is not the out of bound means in the range of 32-bit operating.
		if (reversedNum > math.MaxInt32/10) || (reversedNum < math.MinInt32/10) {
			return 0
		}

		reversedNum = (reversedNum * 10) + digit
		x = x / 10 // next quotient
	}

	return reversedNum
}

func isPalindrome(num int) bool {
	// we don't check the palindrome for negative number
	if num < 0 {
		return false
	}

	reversedNum := reversed(num)
	if num != reversedNum {
		return false
	}

	return true
}

func main() {
	// we don't check the negative number for palindrome
	num := -121
	num1 := 353
	num2 := 1235321
	num3 := 12353123

	fmt.Println(isPalindrome(num))
	fmt.Println(isPalindrome(num1))
	fmt.Println(isPalindrome(num2))
	fmt.Println(isPalindrome(num3))
}
