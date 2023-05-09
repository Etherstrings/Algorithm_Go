package L1605

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	mat := make([][]int, len(rowSum))
	for i, rs := range rowSum {
		mat[i] = make([]int, len(colSum))
		for j, cs := range colSum {
			mat[i][j] = min(rs, cs)
			rs -= mat[i][j]
			colSum[j] -= mat[i][j]
		}
	}
	return mat
}
