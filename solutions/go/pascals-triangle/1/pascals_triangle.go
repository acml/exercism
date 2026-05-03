package pascal

// Triangle computes Pascal's triangle up to a given number of rows.
func Triangle(n int) [][]int {
	res := make([][]int, n)
	for i := 1; i <= n; i++ {
		res[i-1] = row(i)
	}
	return res
}

func row(n int) []int {
	res := make([]int, n)
	res[0] = 1
	res[n-1] = 1
	for k := 1; k <= n/2; k++ {
		res[k] = res[k-1] * (n - k) / k
		res[n-1-k] = res[k]
	}
	return res
}

// recursive solution
// func Triangle(n int) [][]int {
// 	if n == 1 {
// 		return [][]int{{1}}
// 	}

// 	triangle := Triangle(n - 1)
// 	row := make([]int, n)
// 	row[0] = 1
// 	for i := 0; i < len(triangle)-1; i++ {
// 		row[i+1] = triangle[len(triangle)-1][i] + triangle[len(triangle)-1][i+1]
// 	}
// 	row[len(row)-1] = 1

// 	return append(triangle, row)
// }
