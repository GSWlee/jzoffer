package _4

//
//  cuttingRope
//  @Description: 剪绳子(dp)
//  @param n
//  @return int
//
func cuttingRope(n int) int {
	dp := []int{0, 1, 1, 2, 4}
	for i := 5; i <= n; i++ {
		max := i
		for j := 1; j < i; j++ {
			if j*(i-j) > max {
				max = j * (i - j)
			}
			if dp[j]*(i-j) > max {
				max = dp[j] * (i - j)
			}
		}
		dp = append(dp, max)
	}
	return dp[n]
}
