package _0

func dicesProbability(n int) []float64 {
	now, pre := make([]int, 6*n+1), make([]int, 6*n+1)
	for i := 1; i <= 6; i++ {
		now[i] = 1
	}
	for i := 2; i <= n; i++ {
		pre = now
		now = make([]int, 6*n+1)
		for j := i; j <= 6*n; j++ {
			for t := 1; t < j && t <= 6; t++ {
				now[j] += pre[j-t]
			}
		}
	}
	total := 0
	for _, v := range now {
		total += v
	}
	ans := []float64{}
	for _, v := range now {
		if v != 0 {
			temp := float64(v) / float64(total)
			ans = append(ans, temp)
		}
	}
	return ans
}
