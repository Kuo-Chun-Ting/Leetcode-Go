package a5_binets_method

func climbStairs(n int) int {
	q := [][]int{{1, 1}, {1, 0}}
	qN := power(q, n)
	return qN[0][0]
}

func power(matrix [][]int, pow int) [][]int {
	if pow == 0 {
		return [][]int{{1, 0}, {0, 1}}
	}

	if pow == 1 {
		return matrix
	}

	currentPower := 1
	powerOfN := matrix

	for {
		powerOfN = multiply(powerOfN, matrix)
		currentPower++

		if currentPower == pow {
			break
		}
	}
	return powerOfN
}

func multiply(matrix1 [][]int, matrix2 [][]int) [][]int {
	res := [][]int{{0, 0}, {0, 0}}
	res[0][0] = matrix1[0][0]*matrix2[0][0] + matrix1[0][1]*matrix2[1][0]
	res[0][1] = matrix1[0][0]*matrix2[0][1] + matrix1[0][1]*matrix2[1][1]
	res[1][0] = matrix1[1][0]*matrix2[0][0] + matrix1[1][1]*matrix2[1][0]
	res[1][1] = matrix1[1][0]*matrix2[0][1] + matrix1[1][1]*matrix2[1][1]
	return res
}
