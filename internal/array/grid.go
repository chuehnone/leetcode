package array

import (
	"leetcode/internal/math"
)

// Given a 2D grid of size n * m and an integer k. You need to shift the grid k times.
func ShiftGrid(grid [][]int, k int) [][]int {
	if k == 0 {
		return grid
	}

	n := len(grid)
	m := len(grid[0])
	total := n * m
	k = k % total

	rGrid := make([][]int, n)
	for i := 0; i < n; i++ {
		rGrid[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pos := (k + i*m + j) % total
			rGrid[pos/m][pos%m] = grid[i][j]
		}
	}

	return rGrid
}

func MinTimeToVisitAllPoints(points [][]int) int {
	d := 0
	num := len(points)
	for i := 1; i < num; i++ {
		d += math.MaxInt(math.AbsInt(points[i][0]-points[i-1][0]), math.AbsInt(points[i][1]-points[i-1][1]))
	}

	return d
}
