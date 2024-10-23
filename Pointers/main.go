// Decimal Number System Base 10 => (0-9)
// Binary Number System Base 2 => (0/1)
// Octal Number System Base 8 => (0-8)
// HexaDecimal Number System Base 16 => (0-9 & A-F)

// var a = 10  that will be stored in memory
// memory => address location (0x36a6) / HexaDecimal Number System

// Pointers :- Special variables that stores address of other variables.

// AddressOf operator => Ampersand & => To show the actual memory address location
// * => Points Out / Dereference Operator to show the value stored on to that memory address location

package main

import "fmt"

func main() {

	var a = 10
	fmt.Println("The memory address := ", &a)

	var ptr *int = &a // Pointer to store the actual memory address of varaible into ptr varaible
	fmt.Println("The Pointer Reference Address := ", *ptr)
	fmt.Println("The Pointer Actual Memory Address := ", &ptr)

	var parPtr **int = &ptr // Two level pointer
	fmt.Println("The Parent Pointer Reference Address := ", **parPtr)

}
