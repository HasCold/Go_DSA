// Q. 850

// You are given a 2D array of axis-aligned rectangles. Each rectangle[i] = [xi1, yi1, xi2, yi2] denotes the ith rectangle where (xi1, yi1) are the coordinates of the bottom-left corner, and (xi2, yi2) are the coordinates of the top-right corner.

// Calculate the total area covered by all rectangles in the plane. Any area covered by two or more rectangles should only be counted once.

// Return the total area. Since the answer may be too large, return it modulo 109 + 7.

// -----------------------------------------------------------------------------------------------------------
// Modulo Arithmetics :- we find out the modulo arithmetics because sometime our answer  exceeded to the maximum integer limit so we want to know the modular answer

// suppose x % n  ; answer lies between the range of [0 to n - 1]
// 100 % 3  answer lies between the range of [0, 1, 2]

package main

func main() {

}

// ------------------------------- Modulo Properties -----------------------------------

// (x + y) % m = x % m + y % m

// (x - y) % m = x % m - y % m

// (x . y) % m = x % m . y % m

// x = 8, y = 9, m = 3
// 8 % 3 + 9 % 3 = 3 + 0 = 3
