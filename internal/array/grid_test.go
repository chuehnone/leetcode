package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShiftGrid(t *testing.T) {
	assert.Equal(t, ShiftGrid([][]int{{1,2,3},{4,5,6},{7,8,9}}, 1), [][]int{{9,1,2},{3,4,5},{6,7,8}}, "ShiftGrid is incorrect.")
}