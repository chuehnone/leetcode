package main

import (
	"fmt"
	"leetcode/internal/lgame"
)

// LeetCode: https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/
func main() {
	fmt.Println(lgame.GetTicTacToeStatus([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}))
}
