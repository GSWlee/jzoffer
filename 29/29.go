package _9

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	if len(matrix) == 0 {
		return ans
	}
	if len(matrix[0]) == 0 {
		return ans
	}
	start := 0
	col, row := len(matrix[0]), len(matrix)
	for col > 2*start && row > 2*start {
		endx, endy := row-start-1, col-start-1

		for i := start; i <= endy; i++ {
			ans = append(ans, matrix[start][i])
		}
		if start == endx {
			break
		}
		for i := start + 1; i <= endx; i++ {
			ans = append(ans, matrix[i][endy])
		}
		if start == endy {
			break
		}
		for i := endy - 1; i >= start; i-- {
			ans = append(ans, matrix[endx][i])
		}
		for i := endx - 1; i > start; i-- {
			ans = append(ans, matrix[i][start])
		}
		start++
	}
	return ans
}
