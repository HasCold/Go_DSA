package main

import (
	"fmt"
	"log"
	"reflect"
	"sync"
)

func main() {
	fmt.Println("Enter the number for pattern :-")
	var num int
	fmt.Scanln(&num) // pass the reference of actual memory address location of num

	numType := reflect.TypeOf(num)
	if numType.Kind() != reflect.Int { // returns the kind of the type, which can be used to check if it's an int, float, string etc
		log.Fatal("Invalid Integer !")
	}

	syncChan := make(chan struct{})
	syncChan1 := make(chan struct{})
	syncChan2 := make(chan struct{})
	syncChan3 := make(chan struct{})
	syncChan4 := make(chan struct{})
	syncChan5 := make(chan struct{})
	wg := &sync.WaitGroup{}

	wg.Add(7) // I will tell waitGroups that Hey I am adding three goroutines

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		fmt.Println("--------------------- Triangular Patter with increasing number at each line or rows ------------------------")

		for i := 0; i < num; i++ {
			for j := 0; j <= i; j++ { // Ascending Order
				fmt.Print(j+1, " ")
			}
			fmt.Println()
		}

		// close(syncChan)
		syncChan <- struct{}{}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		<-syncChan

		fmt.Println("---------------- Reverse Triangular Patter with increasing number at each line or rows ---------------------")
		for i := 0; i < num; i++ {
			for j := num; j >= i+1; j-- { // Descending Order
				fmt.Print(j, " ")
			}
			fmt.Println()
		}

		syncChan1 <- struct{}{}
		// close(syncChan)
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		<-syncChan1

		fmt.Println("---------------- Floyd's Triangular Pattern ---------------------")
		var counter = 1
		for i := 0; i < num; i++ {
			for j := 0; j <= i; j++ { // Ascending Order
				fmt.Print(counter, " ")
				counter++
			}

			fmt.Println()
		}
		syncChan2 <- struct{}{}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Wait for the third goroutine to finish
		<-syncChan2

		fmt.Println("---------------- Inverted Triangular Pattern ---------------------")
		for i := 0; i < num; i++ {
			for j := 0; j < i; j++ {
				fmt.Print("  ")
			}

			for j := num; j > i; j-- {
				fmt.Print(i+1, " ")
			}
			fmt.Println()
		}

		syncChan3 <- struct{}{} // This will not consume any memory allocation and making it efficient for signaling
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		<-syncChan3

		ch := 'A'

		fmt.Println("---------------- Inverted Triangular Pattern with Alphabets ---------------------")
		for i := 0; i < num; i++ {
			for j := 0; j < i; j++ {
				fmt.Print("  ")
			}

			for j := num; j > i; j-- {
				fmt.Print(string(ch), " ")
			}
			ch++ // A (65) + 1 = B (66) + 1 = C (67) -->> ASCII code
			fmt.Println()
		}

		syncChan4 <- struct{}{}
	}(wg)

	// 				1
	// 			1	2 	1
	// 	 	1	2   3   2   1
	// 1	2	3	4	3	2	1
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		<-syncChan4

		fmt.Println("---------------- Pyramid Pattern ---------------------")

		for i := 0; i < num; i++ {
			for j := num; j > i+1; j-- {
				fmt.Print("  ")
			}

			for j := 0; j <= i; j++ {
				fmt.Print(j+1, " ")
			}

			for j := i; j > 0; j-- {
				fmt.Print(j, " ")
			}

			fmt.Println()
		}
		syncChan5 <- struct{}{}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		<-syncChan5
		fmt.Println("---------------- Hollow Diamond Pattern ---------------------")

		for i := 0; i < num; i++ {
			for j := num; j > i+1; j-- {
				fmt.Print(" ")
			}

			for j := i; j == i; j++ {
				fmt.Print("*")
			}

			for j := 1; j <= ((2 * i) - 1); j++ {
				fmt.Print(" ")
			}

			if i != 0 {
				for j := i; j == i; j++ {
					fmt.Print("*")
				}
			}

			// for j := 0; j >= nu ;j++ {
			// 	fmt.Print(" ")
			// }

			fmt.Println()
		}

	}(wg)

	wg.Wait() // ensure that main goroutine or function will wait the remaining goroutines execution to be compeleted
}
