package _1

func reversePairs(nums []int) int {
	ans := 0
	var merge func([]int) []int
	merge = func(num []int) []int {
		if len(num) <= 1 {
			return num
		} else if len(num) == 2 {
			if num[0] > num[1] {
				ans++
				return []int{num[1], num[0]}
			}
			return num
		} else {
			left, right := merge(num[:len(num)/2]), merge(num[len(num)/2:])
			if len(left) == 0 {
				return right
			} else if len(right) == 0 {
				return left
			} else {
				an := make([]int, len(num))
				ptr1, ptr2 := len(left)-1, len(right)-1
				temp := len(num) - 1
				for ptr1 >= 0 || ptr2 >= 0 {
					if ptr1 == -1 {
						an[temp] = right[ptr2]
						ptr2--
						temp--
					} else if ptr2 == -1 {
						an[temp] = left[ptr1]
						ptr1--
						temp--
					} else {
						if right[ptr2] >= left[ptr1] {
							an[temp] = right[ptr2]
							ptr2--
							temp--
						} else {
							an[temp] = left[ptr1]
							ptr1--
							temp--
							ans = ans + (ptr2 + 1)
						}
					}
				}
				return an
			}
		}
	}
	merge(nums)
	return ans
}
