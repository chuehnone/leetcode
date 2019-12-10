package lstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalancedStringSplit(t *testing.T) {
	assert.Equal(t, 4, BalancedStringSplit("RLRRLLRLRL"), "BalancedStringSplit is incorrect.")
	assert.Equal(t, 3, BalancedStringSplit("RLLLLRRRLR"), "BalancedStringSplit is incorrect.")
	assert.Equal(t, 1, BalancedStringSplit("LLLLRRRR"), "BalancedStringSplit is incorrect.")
	assert.Equal(t, 2, BalancedStringSplit("RLRRRLLRLL"), "BalancedStringSplit is incorrect.")
}

func TestToHexspeak(t *testing.T) {
	assert.Equal(t, "IOI", ToHexspeak("257"))
	assert.Equal(t, "ERROR", ToHexspeak("3"))
}

func TestMaxNumberOfBalloons(t *testing.T) {
	assert.Equal(t, 1, MaxNumberOfBalloons("nlaebolko"))
	assert.Equal(t, 2, MaxNumberOfBalloons("loonbalxballpoon"))
	assert.Equal(t, 0, MaxNumberOfBalloons("leetcode"))
}

func TestSubtractProductAndSum(t *testing.T) {
	assert.Equal(t, 15, SubtractProductAndSum(234))
	assert.Equal(t, 21, SubtractProductAndSum(4421))
}