package main

import (
	"fmt"
	"leetcode/internal/larray"
)

// LeetCode: https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/
func main() {
	fmt.Println(larray.OddCells(2, 3, [][]int{{0, 1}, {1, 1}}))
}
