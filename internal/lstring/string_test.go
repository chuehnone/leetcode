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