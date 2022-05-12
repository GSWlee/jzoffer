package _1

func exchange(nums []int) []int {
	start, end := 0, len(nums)-1
	for start < end {
		if nums[start]%2 == 1 {
			start++
			continue
		}
		if nums[end]%2 == 0 {
			end--
			continue
		}
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--

	}
	return nums
}
