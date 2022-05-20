package _6

func translateNum(num int) int {
	flag := []int{}
	for num > 0 {
		flag = append(flag, num%10)
		num = (num - num%10) / 10
	}
	dp := make([]int, len(flag))
	dp = append(dp, 1)
	dp = append(dp, 1)
	for j := len(flag) - 1; j >= 0; j-- {
		if j+1 < len(flag) {
			if flag[j+1] == 1 || (flag[j+1] == 2 && flag[j] < 6) {
				dp[j] += dp[j+2]
			}
		}
		dp[j] += dp[j+1]
	}
	return dp[0]
}
