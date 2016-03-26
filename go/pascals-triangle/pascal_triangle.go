package pascal

func nextRow(prevRow []int) []int {
	current := make([]int, 0, len(prevRow)+1)

	leftmost := 1
	current = append(current, leftmost)

	for i := 1; i < len(prevRow); i++ {
		current = append(current, prevRow[i-1]+prevRow[i])
	}

	rightmost := 1
	current = append(current, rightmost)
	return current
}

// Triangle takes a number indicating the pascal's triangle depth, and returns the pascal's triangle according to that depth
func Triangle(depth int) [][]int {
	triangle := [][]int{}
	if depth <= 0 {
		return triangle
	}
	triangle = append(triangle, []int{1})
	for i := 1; i < depth; i++ {
		triangle = append(triangle, nextRow(triangle[i-1]))
	}
	return triangle
}
