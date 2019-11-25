package main

import (
	"fmt"
	"leetcode/internal/array"
)

// LeetCode: https://leetcode.com/problems/minimum-time-visiting-all-points/
// On a plane there are n points with integer coordinates points[i] = [xi, yi]. Your task is to find the minimum time in seconds to visit all points.
func main() {
	d := array.MinTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}})
	fmt.Println(d)
}
