package a1_dynamic_programming

func generate(numRows int) [][]int {
	pascalNumber := [][]int{}

	if numRows >= 1 {
		pascalNumber = append(pascalNumber, []int{1})
	}

	if numRows >= 2 {
		pascalNumber = append(pascalNumber, []int{1, 1})
	}

	if numRows >= 3 {
		for i := 2; i < numRows; i++ {
			pascalNumber = append(pascalNumber, []int{}) // init
			pascalNumber[i] = append(pascalNumber[i], 1) // first item must be 1

			lastLayerNumber := pascalNumber[i-1]
			for k := 0; k < len(lastLayerNumber)-1; k++ {
				pascalNumber[i] = append(pascalNumber[i], lastLayerNumber[k]+lastLayerNumber[k+1])
			}

			pascalNumber[i] = append(pascalNumber[i], 1) // last item must be 1
		}
	}

	return pascalNumber
}
