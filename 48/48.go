package _8

func lengthOfLongestSubstring(s string) int {
	dp := []int{1}
	dict := map[rune]int{}
	temp := 0
	max := 0
	for k, v := range s {
		if v, ok := dict[v]; ok {
			if k-v <= temp {
				dp = append(dp, k, v)
				temp = k - v
			} else {
				dp = append(dp, dp[len(dp)-1]+1)
				temp++
			}
		} else {
			dp = append(dp, dp[len(dp)-1]+1)
			temp++
		}
		dict[v] = k
		if temp > max {
			max = temp
		}
	}
	return max
}
