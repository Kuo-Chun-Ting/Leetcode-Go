package a1_rotate_groups_of_four_cells

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}
