package main

import (
	"fmt"
	"leetcode/internal/larray"
)

// LeetCode: https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/
func main() {
	fmt.Println(larray.FindSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}
