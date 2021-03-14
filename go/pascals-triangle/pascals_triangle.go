package pascal

// Triangle computes Pascal's triangle up to a given number of rows.
func Triangle(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	triangle := Triangle(n - 1)
	row := make([]int, n)
	row[0] = 1
	for i := 0; i < len(triangle)-1; i++ {
		row[i+1] = triangle[len(triangle)-1][i] + triangle[len(triangle)-1][i+1]
	}
	row[len(row)-1] = 1

	return append(triangle, row)
}
