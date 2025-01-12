package a1_hashset

func isValidSudoku(board [][]byte) bool {
	rowHashMap := make(map[int]map[byte]bool, 9)
	columnHashMap := make(map[int]map[byte]bool, 9)
	boxHashMap := make(map[int]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rowHashMap[i] = make(map[byte]bool)
		columnHashMap[i] = make(map[byte]bool)
		boxHashMap[i] = make(map[byte]bool)
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num := board[row][col]

			if num == '.' {
				continue
			}

			if rowHashMap[row][num] {
				return false
			}
			rowHashMap[row][num] = true

			if columnHashMap[col][num] {
				return false
			}
			columnHashMap[col][num] = true

			boxIdx := (row/3)*3 + (col / 3)
			if boxHashMap[boxIdx][num] {
				return false
			}
			boxHashMap[boxIdx][num] = true
		}
	}

	return true
}
