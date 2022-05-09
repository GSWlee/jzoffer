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

//
//  cuttingRope
//  @Description: 剪绳子(贪心)
//  @param n
//  @return int
//
func cuttingRope(n int) int {
	length := 1
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	for n >= 5 {
		length = length * 3 % 1000000007
		n = n - 3
	}
	return length * n % 1000000007
}
