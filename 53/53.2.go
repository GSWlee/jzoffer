package _3

func missingNumber(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == mid {
			start = mid + 1
		} else {
			if mid == 0 || (nums[mid-1] == (mid - 1)) {
				if mid == 0 {
					return 0
				} else {
					return mid
				}
			} else {
				end = mid - 1
			}
		}
	}
	if nums[start] == start {
		return start + 1
	}
	return nums[start] - 1
}
