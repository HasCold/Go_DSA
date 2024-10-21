// Constraints:

// 1 <= prices.length <= 105
// 0 <= prices[i] <= 104

// To calculate log 2(10), you can use the change of base formula:
// logb​ (a)= logk​(a) / log k​(b)​

// where 𝑘 can be any positive number (commonly 10 or 𝑒). Using base 10:
// log⁡2(10) = log⁡10(10) / log⁡10(2)

// Question :-
// 1. Buy and Selling day must be different;
// 2. Choosing a different day in the future to sell that stock.
// 3. Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

package main

import "fmt"

func maxProfit(prices []int) int { // O(n) Time Complexity
	var bestBuy = prices[0]
	var maximumProfit = 0

	for i := 0; i < len(prices); i++ {
		if prices[i] > bestBuy {
			maximumProfit = max(maximumProfit, prices[i]-bestBuy)
		}

		bestBuy = min(bestBuy, prices[i])
	}

	return maximumProfit
}

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println("The maximum profit we got from stocks", maxProfit(prices))

}
