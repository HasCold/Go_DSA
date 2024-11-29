package main

import "fmt"

// GCD  -->> Greatest Common Divisor
// HCF  -->> Highest Common Factor

// Better Approach available is Euclid's Algorithm

// gcd(a,b) => gcd(a-b, b)  ; a > b
// gcd(a,b) => gcd(a, b-a)  ; a < b

// gcd(20, 28)  => gcd(20, 8) => gcd(12, 8) => gcd(4, 8) => gcd(4, 4) => gcd(0, 4)  ==>> 4 GCD / HCF of 20, 28

func euclidAlgo(a, b int) int {

	for a > 0 && b > 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}

	if a == 0 {
		return b
	}

	return a
}

func gcdRec(a, b int) int { // Recursive Approach
	if b == 0 {
		return a
	}

	return gcdRec(b, a%b)
}

func lcm(a, b int) int {
	gcd := gcdRec(a, b)
	return (a * b) / gcd
}

func main() {
	a, b := 20, 28

	fmt.Println("The LCM", lcm(a, b))

	fmt.Println(euclidAlgo(a, b))
	fmt.Println(gcdRec(a, b))
}
