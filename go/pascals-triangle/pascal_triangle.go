package pascal

func pascal(row, col int) int {
	if col == 0 || row == col {
		return 1
	}
	return pascal(row-1, col-1) + pascal(row-1, col)
}

// Triangle takes a number indicating the pascal's triangle depth, and returns the pascal's triangle according to that depth
func Triangle(depth int) [][]int {
	triangle := make([][]int, depth)
	for row := 0; row < depth; row++ {
		for col := 0; col <= row; col++ {
			triangle[row] = append(triangle[row], pascal(row, col))
		}
	}
	return triangle
}
