package larray

import (
	"leetcode/internal/math"
	"sort"
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
		if dX*tY != tX*dY {
			return false
		}
	}

	return true
}

func MinCostToMoveChips(chips []int) int {
	even := 0
	odd := 0

	for _, v := range chips {
		if v&1 == 1 {
			odd++
		} else {
			even++
		}
	}

	return math.MinInt(even, odd)
}

func UniqueOccurrences(arr []int) bool {
	countMap := make(map[int]int)

	for _, v := range arr {
		countMap[v]++
	}

	uniqueMap := make(map[int]bool)
	for _, count := range countMap {
		if _, exist := uniqueMap[count]; exist {
			return false
		}
		uniqueMap[count] = true
	}

	return true
}

func MinimumAbsDifference(arr []int) [][]int {
	arrLen := len(arr)

	sort.Ints(arr)

	min := arr[1] - arr[0]
	ans := [][]int{{arr[0], arr[1]}}
	for i := 2; i < arrLen; i++ {
		dis := arr[i] - arr[i-1]
		if min > dis {
			min = dis
			ans = [][]int{{arr[i-1], arr[i]}}
		} else if min == dis {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}

	return ans
}

// https://www.wikiwand.com/en/Determination_of_the_day_of_the_week#/Implementation-dependent_methods
func DayOfTheWeek(day int, month int, year int) string {
	if month < 3 {
		day = day + year
		year--
	} else {
		day += year - 2
	}

	dayOfWeek := (23*month/9 + day + 4 + year/4 - year/100 + year/400) % 7
	weekTable := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	return weekTable[dayOfWeek]
}
