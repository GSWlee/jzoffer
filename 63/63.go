package _3

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, max := prices[0], 0
	for _, v := range prices {
		if v < min {
			min = v
		}
		if v-min > max {
			max = v - min
		}
	}
	return max
}
