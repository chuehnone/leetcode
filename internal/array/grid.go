package array

// Given a 2D grid of size n * m and an integer k. You need to shift the grid k times.
func ShiftGrid(grid [][]int, k int) [][]int {
	if k == 0 {
		return grid
	}

	n := len(grid)
	m := len(grid[0])
	temp := grid[n-1][m-1]

	rGrid := make([][]int, n)
	for i := 0; i < n; i++ {
		rGrid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			rGrid[i][j] = temp
			temp = grid[i][j]
		}
	}

	return ShiftGrid(rGrid, k-1)
}
