package _7

func findContinuousSequence(target int) [][]int {
	ans := [][]int{}
	sum, start, end := 1, 1, 1
	for end < target {
		if sum < target {
			end++
			sum += end
		} else if sum > target {
			sum -= start
			start++
		} else {
			tmp := []int{}
			for i := start; i <= end; i++ {
				tmp = append(tmp, i)
			}
			ans = append(ans, tmp)
			end++
			sum += end
		}
	}
	return ans
}
