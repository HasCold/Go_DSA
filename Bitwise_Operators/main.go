package main

import (
	"fmt"
	"unsafe"
)

// Bitwise Operators :- Are similar to the Logical AND , OR and XOR operator.
// 1. Bitwise &
// 2. Bitwise |
// 1. Bitwise ^

// Bitwise AND Operator :-
//  0(bit) & 0(bit)  = 0 bit
//  0(bit) & 1(bit)  = 0 bit
//  1(bit) & 0(bit)  = 0 bit
//  1(bit) & 1(bit)  = 1 bit

// Bitwise OR Operator :-
//  0(bit) | 0(bit)  = 0 bit		Taking bitwise | operator between two bits of integer 4 and 8 :-
//  1(bit) | 0(bit)  = 1 bit			4 = 0100
//  0(bit) | 1(bit)  = 1 bit			8 = 1000
//  1(bit) | 1(bit)  = 1 bit			   (1100) = (12)

//
// Bitwise XOR ^ (Exclusive OR) Operator :-
// same bit  -->> 0
// diff bit -->> 1

//  0(bit) ^ 0(bit)  = 0 bit  		Taking bitwise ^ operator between two bits of integer 4 and 8 :-
//  0(bit) ^ 1(bit)  = 1 bit			4 = 0100
//  1(bit) ^ 0(bit)  = 1 bit			8 = 1000
//  1(bit) ^ 1(bit)  = 0 bit			   (1100) = (12)

//
// Bitwise Left-Shift Operator :- <<  (Points to the left)
//  n = 4 (100)
//  n << 1  -->> 1000 (8) ; bit shift towards the left side

// Left-Shift operator internal memory operation :- ans = a * 2^b /  4 * 2^1 = 8

//
// Bitwise Right-Shift Operator :- >> (Points to the right)
//  n = 4 (100)
//  n >> 1  -->> 10 (2) ; bit shift towards the right side and when the bit move towards the right-side so the bit(trailing value) at the end will drop from the memory.

// Right-Shift operator internal memory operation :- ans = a/2^b / 8/2^1

func isPowOf2(p int) {
	pow := 1 // 2^0 = 1

	for i := 0; i <= p; i++ {
		pow *= 2
		if pow == p {
			fmt.Println("The Power of 2 :-", p)
			break
		}
	}

}

func main() {

	var a int
	var b float64
	var c bool
	var unsigned uint

	fmt.Println("The bitwise &, |, ^, << and >> operator :-", a<<1)

	var p = 125
	isPowOf2(p)

	// Data-Modifiers :-
	// unsafe.Sizeof() returns the size of a variable in bytes, similar to C++'s sizeof.

	// 1 byte = 8 bits
	// int has given the 4-bytes memory allocation in memory :- 4 bytes = 32 bits => -2^31 to +2^31 - 1
	// if int64 has given the 8-bytes memory allocation in memory :- 8 bytes = 64 bits => -2^63 to +2^63 - 1

	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(unsigned))

}

// Bitwise AND (`&`), OR (`|`), and XOR (`^`) operators are similar in concept to their logical counterparts (`&&`, `||`, `^`), but they operate at the bit level rather than on Boolean values. Here's a comparison:

// 1. **Bitwise Operators:**
//    - **AND (`&`)**: Performs a bitwise AND operation between corresponding bits of two integers. The result is `1` only if both corresponding bits of operands are `1`; otherwise, it's `0`.
//    - **OR (`|`)**: Performs a bitwise OR operation between corresponding bits of two integers. The result is `1` if at least one of the corresponding bits of operands is `1`.
//    - **XOR (`^`)**: Performs a bitwise XOR (exclusive OR) operation between corresponding bits of two integers. The result is `1` if the corresponding bits of operands are different; otherwise, it's `0`.

// 2. **Logical Operators:**
//    - **Logical AND (`&&`)**: Returns `true` if both operands are `true`; otherwise, it returns `false`. It short-circuits if the first operand evaluates to `false`.
//    - **Logical OR (`||`)**: Returns `true` if at least one operand is `true`; otherwise, it returns `false`. It short-circuits if the first operand evaluates to `true`.
//    - **Logical XOR (`^`)**: Returns `true` if exactly one operand is `true` and the other is `false`; otherwise, it returns `false`.

// **Key Differences:**
// - Bitwise operators work on individual bits of integers.
// - Logical operators work with Boolean values (`true` or `false`).
// - Bitwise operators always evaluate both operands.
// - Logical operators may short-circuit, meaning they might not evaluate the second operand if the result can be determined from the first operand.

// In summary, while bitwise AND, OR, and XOR operators share some conceptual similarities with their logical counterparts, they operate on a more fundamental level of individual bits within integer representations.
