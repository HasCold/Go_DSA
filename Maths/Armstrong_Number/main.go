package main

import (
	"fmt"
	"math"
)

// Armstrong Number is a number that is equal to the sum of cubes of its digits.
// eg. :- 153, 371, 1

// 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153

// 3^3 + 7^3 + 1^3  = 27 + 343 + 1 = 371

func isArmStrongNum(n int) bool {
	copyOfN := n // Base 10 (decimal number system)
	sumOfCubes := 0

	for n != 0 {
		digit := n % 10 // Get a single digit
		sumOfCubes += int(math.Pow(float64(digit), 3))
		n = n / 10 // Update the n value
		// sumOfCubes += (digit*digit*digit)
	}

	return sumOfCubes == copyOfN
}

func main() {

	n := 153

	if isArmStrongNum(n) {
		fmt.Println(n, "is an Armstrong Number")
	} else {
		fmt.Println(n, "is not an Armstrong Number")
	}

}
