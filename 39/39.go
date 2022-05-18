package _9

func majorityElement(nums []int) int {
	target, num := 0, 0
	for _, v := range nums {
		if num == 0 {
			target = v
			num++
		} else {
			if target == v {
				num++
			} else {
				num--
			}
		}
	}
	return target
}
