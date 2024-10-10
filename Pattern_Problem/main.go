package main

import (
	"fmt"
	"log"
	"reflect"
	"sync"
)

// Pattern Problem is very helpful in the Dynamic Programming (DP)

// Columns  (That print the inner loop)
//   ^
//   |
// 1 2 3 4 	-> Rows (That print the outer loop)
// 1 2 3 4	-> Rows (That print the outer loop)
// 1 2 3 4	-> Rows (That print the outer loop)
// 1 2 3 4	-> Rows (That print the outer loop)

// Outer loop basically tells you the no. of line that you want to print.
// Inner loop print that thing that you want at the specific line.

func main() {

	fmt.Println("Enter the number")
	var num int
	fmt.Scanln(&num)

	numType := reflect.TypeOf(num)

	if numType.Kind() != reflect.Int { // returns the kind of the type, which can be used to check if it's an int, float, string etc
		log.Fatal("Invalid Integer")
	}

	var wg sync.WaitGroup

	// Channels to synchronize the order of execution
	squareDone := make(chan bool)
	starDone := make(chan bool)

	wg.Add(3) // I will tell waitGroups that Hey I am adding three goroutines

	go func() {
		fmt.Println("-------------------- Square Pattern Start -----------------")
		for i := 1; i <= num; i++ { // Outer loop
			for j := 1; j <= num; j++ { // Inner loop
				fmt.Print(j, " ") // Print the numbers in the same line
			}
			fmt.Println() // Print a new line after each row
		}
		wg.Done()          // And after completing operation or task then Done
		squareDone <- true // Notify the next goroutine to start
	}()

	// Star Pattern
	go func() {
		<-squareDone
		fmt.Println("-------------------- Star Pattern Start --------------------")
		for i := 1; i <= num; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print(" * ") // Print the characters at the same line
			}
			fmt.Println() // Print a new line after each row
		}
		wg.Done()
		starDone <- true
	}()

	// In mapping there is a hashing data structure / hash-map
	alpha := map[int]string{1: "A", 2: "B", 3: "C", 4: "D"}

	// ABCD Patter
	go func() {
		<-starDone
		fmt.Println("-------------------- Character Pattern Start --------------------")

		for i := 1; i <= num; i++ { // Outer loop prints the no. of line
			for j := 1; j <= i; j++ { // Inner loop prints that you want at the row
				if _, ok := alpha[j]; !ok {
					log.Fatal("Alphabet not present : ", ok)
				}
				fmt.Print(alpha[j], " ")
			}
			fmt.Println() // Print a new line after each row
		}

		wg.Done()
	}()

	wg.Wait() // ensures that the main goroutine waits for all three patterns to complete.
}

// What Problem Do WaitGroups Solve?

// In Go, when the main() function returns, the program terminates, and all running goroutines are killed, even if they haven't finished their execution. If you're launching several goroutines but want the main goroutine to wait for them to complete their work, WaitGroups are essential. Without a WaitGroup, your goroutines might not complete their tasks because the main() function could finish execution and terminate the program prematurely.
