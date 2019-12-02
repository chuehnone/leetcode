package lgame

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTicTacToeStatus(t *testing.T) {
	assert.Equal(t, "A", GetTicTacToeStatus([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}))
	assert.Equal(t, "B", GetTicTacToeStatus([][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}))
	assert.Equal(t, "Draw", GetTicTacToeStatus([][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}))
	assert.Equal(t, "Pending", GetTicTacToeStatus([][]int{{0, 0}, {1, 1}}))
}
