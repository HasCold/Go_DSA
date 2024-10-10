package main

import (
	"errors"
	"fmt"
	"log"
)

// Prime number := Are those which include only into the table 1 and table of themselves
// Prime number cannot be present in proceeding bigger number of tables like 7 is the factor present in 1 and 7 table but not in 8, 9, 10... so on.

func isPrime(isPrimeNum int) (res int, err error) {
	initialCheck := 2 // It is a short hand method to declare a variable and the compiler will detect or judges their data-type at the compile time.

	if isPrimeNum <= 1 {
		return 0, errors.New("Enter the number greater than 1")
	}

	if initialCheck < (isPrimeNum - 1) {
		for {
			if isPrimeNum%initialCheck != 0 {

				initialCheck++
				if initialCheck == (isPrimeNum - 1) {
					fmt.Printf("The number : %v is a prime number \n", isPrimeNum)
					return isPrimeNum, nil
				}

			} else if isPrimeNum%initialCheck == 0 {
				fmt.Printf("The number : %v is not the prime number \n", isPrimeNum)
				return isPrimeNum, nil
			}
		}
	} else {
		fmt.Printf("The number : %v is a prime number \n", isPrimeNum)
		return isPrimeNum, nil
	}
}

func main() {

	fmt.Println("Enter the number :- ")
	var num int
	fmt.Scanln(&num)

	res, err := isPrime(num)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("The Result is %v and their data-type : %T \n", res, res)
	}
}
