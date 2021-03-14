package pascal

func Triangle(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	} else if n == 2 {
		return append(Triangle(1), []int{1, 1})
	}

	pres := Triangle(n - 1)
	res := []int{1}
	for i := 0; i < len(pres)-1; i++ {
		res = append(res, pres[len(pres)-1][i]+pres[len(pres)-1][i+1])
	}
	res = append(res, 1)
	return append(Triangle(n-1), res)
}
