package main

import "fmt"

// Quick Summary:
// Even: Divisible by 2 (e.g., 2, 4, 6).
// Odd: Not divisible by 2 (e.g., 1, 3, 5).
// Prime: Has only two divisors, 1 and itself (e.g., 2, 3, 5).
// Composite: Has more than two divisors (e.g., 4, 6, 9).

func isPrime(n int) string {
	for i := 2; i*i <= n; i++ {
		// fmt.Println(i * i)
		if n%i == 0 {
			return "Non Prime"
		}
	}

	return "Prime"
}

// -------------------------- Sieve of Eratosthenes ------------------------------
// If we have a range from 1 to N numbers and we have to find those numbers which are prime and non-prime
// Officially the prime number start from 2 not with the 1 becuase 1 is not the prime or non-prime number

// We can consider firstly that all the numbers are prime numbers we are going to false them on slices

func sieveOfEratosthenes(n int) int {

	length := n + 1
	isPrime := make([]bool, length)

	for i := range isPrime {
		isPrime[i] = true
	}

	count := 0

	for i := 2; i < length; i++ {
		if isPrime[i] {
			fmt.Println(i, isPrime[i])
			count++

			for j := i * 2; j < n; j = j + i {
				isPrime[j] = false
			}
		}

	}

	return count
}

func main() {
	n := 79
	n1 := 50

	fmt.Println(sieveOfEratosthenes(n1))
	fmt.Println(isPrime(n))
}
