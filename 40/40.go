package _0

func getLeastNumbers(arr []int, k int) []int {
	var fz func(s, e int) int
	fz = func(s, e int) int {
		index := s
		i, j := s, e
		for i <= j {
			if arr[i] <= arr[index] {
				i++
				continue
			} else {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
				j--
			}
		}
		temp := arr[index]
		arr[index] = arr[j]
		arr[j] = temp
		return i - 1
	}
	if k == 0 {
		return []int{}
	}
	index, s, e := -1, 0, len(arr)-1
	for index != k-1 {
		index = fz(s, e)
		if index < k-1 {
			s = index + 1
		}
		if index > k-1 {
			e = index - 1
		}
	}
	return arr[:index+1]
}
