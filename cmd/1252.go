package main

import (
	"fmt"
	"leetcode/internal/array"
)

// LeetCode: https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/
func main() {
	fmt.Println(array.OddCells(2, 3, [][]int{{0, 1}, {1, 1}}))
}
