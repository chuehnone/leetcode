package main

import (
	"fmt"
	"leetcode/internal/larray"
)

// LeetCode: https://leetcode.com/problems/shift-2d-grid/
// Given a 2D grid of size n * m and an integer k. You need to shift the grid k times.
func main() {
	r := larray.ShiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1)
	fmt.Printf("%v", r)
}
