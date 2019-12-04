package larray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShiftGrid(t *testing.T) {
	assert.Equal(t, [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}}, ShiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1), "ShiftGrid is incorrect.")
	assert.Equal(t, [][]int{{1}}, ShiftGrid([][]int{{1}}, 100), "ShiftGrid is incorrect.")
	assert.Equal(t, [][]int{{6}, {5}, {1}, {2}, {3}, {4}, {7}}, ShiftGrid([][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}}, 23), "ShiftGrid is incorrect.")
}

func TestMinTimeToVisitAllPoints(t *testing.T) {
	assert.Equal(t, 7, MinTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}), "MinTimeToVisitAllPoints is incorrect.")
}

func TestOddCells(t *testing.T) {
	assert.Equal(t, 6, OddCells(2, 3, [][]int{{0, 1}, {1, 1}}), "OddCells is incorrect.")
	assert.Equal(t, 0, OddCells(2, 2, [][]int{{1, 1}, {0, 0}}), "OddCells is incorrect.")
}

func TestCheckStraightLine(t *testing.T) {
	assert.Equal(t, true, CheckStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}), "CheckStraightLine is incorrect.")
	assert.Equal(t, false, CheckStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}), "CheckStraightLine is incorrect.")
	assert.Equal(t, true, CheckStraightLine([][]int{{-3, -2}, {-1, -2}, {2, -2}, {-2, -2}, {0, -2}}), "CheckStraightLine is incorrect.")
	assert.Equal(t, true, CheckStraightLine([][]int{{-2, 12}, {2, -8}, {6, -28}, {-10, 52}, {-7, 37}, {4, -18}, {7, -33}, {1, -3}, {-1, 7}, {8, -38}}), "CheckStraightLine is incorrect.")
}

func TestMinCostToMoveChips(t *testing.T) {
	assert.Equal(t, 1, MinCostToMoveChips([]int{1, 2, 3}), "MinCostToMoveChips is incorrect.")
	assert.Equal(t, 2, MinCostToMoveChips([]int{2, 2, 2, 3, 3}), "MinCostToMoveChips is incorrect.")
}

func TestUniqueOccurrences(t *testing.T) {
	assert.Equal(t, true, UniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	assert.Equal(t, false, UniqueOccurrences([]int{1, 2}))
	assert.Equal(t, true, UniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}

func TestMinimumAbsDifference(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2}, {2, 3}, {3, 4}}, MinimumAbsDifference([]int{4, 2, 1, 3}))
}
