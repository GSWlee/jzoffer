package _1

import "sort"

func isStraight(nums []int) bool {
	sort.Ints(nums)
	nums0, numsq := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums0++
			continue
		}
		if i > 0 {
			if nums[i] == nums[i-1] {
				return false
			}
			if nums[i-1] != 0 {
				numsq += (nums[i] - nums[i-1] - 1)
			}
		}
	}
	if nums0 < numsq {
		return false
	}
	return true
}
