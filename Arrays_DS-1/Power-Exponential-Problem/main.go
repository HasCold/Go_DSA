// level - Medium

// Compute x^n
// Question :- Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
// Constraints :-
// -100.0 < x < 100.0
// -2^31 <= n <= 2^31 - 1

// We know that our algorithm or code must perform in 1 second = ~10^8 no. of operations
// so n = 2^31 no. of operations > 10^8 (TLE hit : Time Limit Exceeded)
// O(n) Linear T.C. will work

// We can solve this question by the binary exponentiation
//  Binary exponentiation Approach :-
// we run a loop on the binary form of the power ; 3^5  so loop will run on the binary form of 5 = 101
// So according to this we got the time complexity reduces to the O(logn)

// Suppose we have a decimal number n -->> they have atmost binaryForm digit equal to => log base 2 of n + 1
// n = 8  ==>> 1000

// The logarithm base 2 of 8 is:
// log_2(8) = 3 (since 2^3 = 8)
// Now, adding 1 to it:
// 3 + 1 = 4
// So, (log_2(8) + 1 = 4).
// So, 8 has a 4 digits in the binary form.

// 3^5 => Binary exponentiation of power 5 = 101
//  	1 * 3^4 + 0 * 3^2 + 1 * 3^1  (0 will be ignored)
//  	3^4 *  3^1
//  ans = 3^5

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func decToBinary(n int) int {

// 	var rem int
// 	var ans int
// 	pow := 1

// 	for n > 0 {
// 		rem = n % 2 // modulus : remainder
// 		ans = rem * pow
// 		pow *= 10
// 		n /= 2 // Quotient
// 	}
// 	fmt.Printf("The Binary Form is %v \n", ans)

// 	return ans
// }

func myPow(x float64, n int) float64 { // Time Complexity => O(logn)

	binaryForm := int(n)
	ans := 1.0

	// Edge Cases
	if n == 0 {
		return 1
	}

	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	if x == -1 && n%2 == 0 {
		return 1
	}

	if x == -1 && n%2 != 0 {
		return -1
	}

	if n < 0 {
		x = 1 / x
		binaryForm = -binaryForm
	}

	for binaryForm > 0 {
		if binaryForm%2 == 1 {
			ans *= x
		}
		x *= x
		binaryForm /= 2

	}

	return ans
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	// Taking input for base (x)
	fmt.Println("Enter the base (x) :- ")
	baseInput, _ := reader.ReadString('\n')
	baseStr := strings.TrimSpace(baseInput)
	x, err := strconv.ParseFloat(baseStr, 64)
	if err != nil {
		fmt.Println("Invalid Input for base")
	}

	// Taking input for exponent (n)
	fmt.Println("Enter the exponent (n) :- ")
	expInp, _ := reader.ReadString('\n')
	expTrim := strings.TrimSpace(expInp)
	n, err := strconv.Atoi(expTrim)
	if err != nil {
		fmt.Println("Invalid Input for exponent")
	}

	// Binary Exponential Approach
	fmt.Printf("The base %v of exponent %v is :- %v", x, n, myPow(x, n))

}

// Binary Number System for 0 - 10
// 0 = 00
// 1 = 01
// 2 = 10
// 3 = 11
// 4 = 100
// 5 = 101
// 6 = 110
// 7 = 111
// 8 = 1000
// 9 = 1001
// 10 = 1010
// 11 = 1011
// 12 = 1100
