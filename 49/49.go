package _9

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0
	dp = append(dp, 1)
	for i := 1; i < n; i++ {
		x1, x2, x3 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		x := min(min(x1, x2), x3)
		if x == x1 {
			p2++
		}
		if x == x2 {
			p3++
		}
		if x == x3 {
			p5++
		}
		dp[i] = x
	}
	return dp[n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
