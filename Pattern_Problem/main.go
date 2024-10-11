package main

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"
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
	// struct{} is used because it consumes no memory. It’s a way to send a signal without any actual data.
	syncChan := make(chan struct{})

	wg.Add(5) // I will tell waitGroups that Hey I am adding three goroutines

	go func() {
		defer wg.Done() // defer keyword will run the statement after some delay or execution will complete at the end And after completing operation or task then Done

		fmt.Println("-------------------- Square Pattern Start -----------------")
		for i := 1; i <= num; i++ { // Outer loop
			for j := 1; j <= num; j++ { // Inner loop
				fmt.Print(j, " ") // Print the numbers in the same line
			}
			fmt.Println() // Print a new line after each row
		}
		syncChan <- struct{}{} // Notify the next goroutine to start
	}()

	// Star Pattern
	go func() {
		defer wg.Done()
		<-syncChan
		fmt.Println("-------------------- Star Pattern Start --------------------")
		for i := 1; i <= num; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print(" * ") // Print the characters at the same line
			}
			fmt.Println() // Print a new line after each row
		}
		syncChan <- struct{}{}
	}()

	// In mapping there is a hashing data structure / hash-map
	// alpha := map[int]string{1: "A", 2: "B", 3: "C", 4: "D"}

	// ABCD Patter
	go func() {
		defer wg.Done()
		<-syncChan // Send-only Channel
		fmt.Println("-------------------- Character Pattern Start --------------------")

		// Each Character have an ASCII code which means the mathematical representation of any characters.
		// e.g. :- 'A' = 65 ;  if A + 1 = B (66) ; if B + 1 = C (67)

		for i := 1; i <= num; i++ {
			var ch = 'A'
			// fmt.Printf("ASCII Value of %v :- %v \n", ch, rune(ch))

			for j := 1; j <= i; j++ {
				fmt.Print(string(ch), " ")
				ch += 1 // A(65) + 1 = B(66) + 1 = C(67)
			}
			fmt.Println()
		}

		syncChan <- struct{}{} // Recieve-only Channel
	}()

	// Square Pattern with the Continuation of Number
	go func() {
		defer wg.Done()
		<-syncChan // Send-only channel
		fmt.Println("-------------------- Square Pattern With Continuation of Number --------------------")
		var counter = 1

		for i := 1; i <= num; i++ { // Control or Print the no. of lines or no. of rows
			for j := 1; j <= num; j++ { // Print the no. of columns or the thing that you want in a line
				fmt.Print(counter, " ")
				counter++
			}
			fmt.Println()
		}

		syncChan <- struct{}{} // Empty struct doesn't consume any memory, making it efficient for signaling / recieve-only channel
	}()

	// Square Pattern with the Continuation of Number
	go func() {
		defer wg.Done()
		<-syncChan // Send-only Channel
		fmt.Println("-------------------- Square Pattern With Continuation of Alphabets --------------------")

		var counter = 'A' // We can use the rune data-type for converting the character into ASCII code (mathematical representation of any character)

		for i := 0; i < num; i++ {
			for j := 0; j < num; j++ {
				fmt.Print(string(counter), " ")
				counter++
			}
			fmt.Println()
		}

		syncChan <- struct{}{} // Empty struct doesn't consume any memory, making it efficient for signaling / recieve-only channel
	}()

	go func() {
		timeStart := time.Now() // Execution start

		time.Sleep(time.Second)
		fmt.Println("-------------------------------------------------------")
		var ch = 72
		fmt.Println("Converting the character into ASCII code OR conveting integer into character :- ", string(ch))

		elapsedTime := time.Since((timeStart))
		fmt.Println("Execution Time / Elapsed Time :- ", elapsedTime)

		wg.Done()
	}()

	wg.Wait() // ensures that the main goroutine waits for all three patterns to complete.
}

// What Problem Do WaitGroups Solve?

// In Go, when the main() function returns, the program terminates, and all running goroutines are killed, even if they haven't finished their execution. If you're launching several goroutines but want the main goroutine to wait for them to complete their work, WaitGroups are essential. Without a WaitGroup, your goroutines might not complete their tasks because the main() function could finish execution and terminate the program prematurely.

// -------------------------------Channels-------------------------------------------------

// Explanation:
// Single Channel (syncChan): A single channel syncChan of type chan struct{} is used to synchronize all goroutines. Each goroutine waits for its turn by reading from the channel (<-syncChan), and when it finishes its task, it signals the next goroutine by sending a value to the channel (syncChan <- struct{}{}).
// Flow: Each goroutine waits for a signal on the channel and then signals the next one in line by sending an empty struct. This empty struct (struct{}{}) doesn't consume any memory, making it efficient for signaling.
