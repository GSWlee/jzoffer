package _7

func twoSum(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	for end > start {
		if nums[start]+nums[end] < target {
			start++
		} else if nums[start]+nums[end] > target {
			end--
		} else {
			return []int{nums[start], nums[end]}
		}
	}
	return []int{}
}
