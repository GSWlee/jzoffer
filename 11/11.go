package _1

func minArray(numbers []int) int {
	start, end := 0, len(numbers)-1
	for start <= end {
		mid := (start + end) / 2
		if numbers[start] == numbers[end] && numbers[start] == numbers[mid] {
			for i := start + 1; i <= end; i++ {
				if numbers[i] < numbers[i-1] {
					return numbers[i]
				}
			}
			return numbers[start]
		}
		if numbers[start] < numbers[end] {
			return numbers[start]
		}
		if numbers[mid] >= numbers[start] {
			start = mid + 1
			continue
		} else {
			end = mid
		}
	}
	return -1
}
