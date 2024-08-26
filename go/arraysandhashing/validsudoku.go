package arraysandhashing

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]bool, 9)
	cols := make(map[int]map[byte]bool, 9)
	squares := make(map[int]map[byte]bool, 9) // 9 3x3 squares
	const empty byte = 46                     // 46 == "."

	for row := range board {
		rows[row] = make(map[byte]bool)
		for col := range board[row] {
			if _, ok := cols[col]; !ok {
				cols[col] = make(map[byte]bool)
			}

			value := board[row][col]
			if value == empty {
				continue
			}

			// generate unique key for each square
			square := (row / 3) + (col/3)*3
			if _, ok := squares[square]; !ok {
				squares[square] = make(map[byte]bool)
			}

			if rows[row][value] || cols[col][value] || squares[square][value] {
				return false
			}

			rows[row][value] = true
			cols[col][value] = true
			squares[square][value] = true
		}
	}

	return true
}
