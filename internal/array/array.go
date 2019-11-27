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

func OddCells(n int, m int, indices [][]int) int {
	row := make([]int, n)
	col := make([]int, m)
	for _, point := range indices {
		row[point[0]]++
		col[point[1]]++
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (row[i]+col[j])%2 == 1 {
				cnt++
			}
		}
	}

	return cnt
}

func CheckStraightLine(coordinates [][]int) bool {
	dX := coordinates[1][0] - coordinates[0][0]
	dY := coordinates[1][1] - coordinates[0][1]
	num := len(coordinates)

	for i := 2; i < num; i++ {
		tX := coordinates[i][0] - coordinates[0][0]
		tY := coordinates[i][1] - coordinates[0][1]
		if dX * tY != tX * dY {
			return false
		}
	}

	return true
}
