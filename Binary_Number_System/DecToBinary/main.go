package main

import (
	"fmt"
)

// Detailed explanation present in the register

//  n = 5  -->> Convert decimal number system to binary number system ; Base 10 to Base 2
// divisor  dividend
//	 	2 | 5  = 1 (remainder)
// 		2 | 2 (quotient) = 0
// 		2 | 1  = 1
// 	   		0

// Count the number from backward position 101
// 							(1 x 10^2) + (0 x 10^1) + (1 x 10^0)

func decToBinary(decimalNum int) {
	// Convert Decimal N.S to Binary N.S

	var ans int
	pow := 1 // 10^0 = 1, 10^1 = 10, 10^2 = 100
	var rem int

	for decimalNum > 0 {

		rem = decimalNum % 2        //*decimalNum is used to dereference the pointer and access its value for both the % and / operations.
		decimalNum = decimalNum / 2 // (next quotient)

		ans += (rem * pow)
		pow *= 10 // 10^0 = 1 * 10 (10^1) = 10 * 10 = 100 (10^2)
	}

	fmt.Printf("The Binary Number System After Conversion is %v \n", ans)
}

func main() {
	fmt.Println("Enter the decimal number to convert into binary number system :- ")
	var decimalNum int
	fmt.Scan(&decimalNum)

	// wg := sync.WaitGroup{}

	// wg.Add(1)
	for i := 1; i <= decimalNum; i++ {
		decToBinary(i)
	}

	// wg.Wait()
}
