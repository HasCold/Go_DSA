// Question :-

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

package main

import "fmt"

func maxArea(height []int) int { // Time Complexity O(n)

	// Two Pointer Approach
	i := 0
	j := len(height) - 1
	var wd int
	var bd int
	currArea := 0
	maxWater := 0

	for i < j {
		wd = j - i
		bd = min(height[i], height[j])
		currArea = wd * bd
		maxWater = max(maxWater, currArea)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxWater
}

func main() {

	// Area := Length(width)(x-axis) * Breadth(Height)(y-axis)
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	num := maxArea(height)

	fmt.Println("The maximum water container holds :- ", num)

}
