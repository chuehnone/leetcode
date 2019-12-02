package lgame

func GetTicTacToeStatus(moves [][]int) string {
	const aWin = "A"
	const bWin = "B"
	const draw = "Draw"
	const pending = "Pending"

	num := len(moves)

	aMap := make([]int, 8)
	bMap := make([]int, 8)
	for i := 0; i < num; i++ {
		if i&1 == 0 {
			aMap[moves[i][0]]++
			aMap[moves[i][1]+3]++
			if moves[i][0]+moves[i][1] == 2 {
				aMap[6]++
			}
			if moves[i][0] == moves[i][1] {
				aMap[7]++
			}

			if aMap[moves[i][0]] == 3 || aMap[moves[i][1]+3] == 3 || aMap[6] == 3 || aMap[7] == 3 {
				return aWin
			}
		} else {
			bMap[moves[i][0]]++
			bMap[moves[i][1]+3]++
			if moves[i][0]+moves[i][1] == 2 {
				bMap[6]++
			}
			if moves[i][0] == moves[i][1] {
				bMap[7]++
			}

			if bMap[moves[i][0]] == 3 || bMap[moves[i][1]+3] == 3 || bMap[6] == 3 || bMap[7] == 3 {
				return bWin
			}
		}
	}

	if num == 9 {
		return draw
	}
	return pending
}
