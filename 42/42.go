package _2

func maxSubArray(nums []int) int {
	max := nums[0]
	dp := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp = append(dp, dp[i-1]+nums[i])
		} else {
			dp = append(dp, nums[i])
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
