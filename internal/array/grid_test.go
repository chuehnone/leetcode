package array

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