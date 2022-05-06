package main

import (
	"fmt"
)

//
// If_duplicate
//  @Description: 判断数组中是否有重复元素(数组中元素的值小于数组长度)
//  @param nums: 输入数组
//  @return bool：
//

func If_duplicate(nums []int) bool {
	length := len(nums)
	for i := 0; i < length; {
		if nums[i] >= length {
			return true
		} else if nums[i] == i {
			i++
		} else {
			if nums[nums[i]] == nums[i] {
				return true
			} else {
				temp := nums[nums[i]]
				nums[nums[i]] = nums[i]
				nums[i] = temp
			}
		}
	}
	return false
}

//
// Find_duplicate
//  @Description: 查找数组中重复数组(建立辅助数组)
//  @param num: 数组
//  @return int: 重复元素
//
func Find_duplicate(num []int) int {
	nums := append([]int{}, num...)
	length := len(nums)
	for i := 0; i < length; {
		if nums[i] == i {
			i++
		} else {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			} else {
				temp := nums[nums[i]]
				nums[nums[i]] = nums[i]
				nums[i] = temp
			}
		}
	}
	return -1
}

//
// Find_duplicate_v2
//  @Description: 查找数组中重复数组(利用二分查找)
//  @param num
//  @return int
//
func Find_duplicate_v2(nums []int) int {
	start, end := 1, len(nums)-1
	for end >= start {
		middle := (end-start)>>1 + start
		counts := count(nums, start, middle)
		if start == end {
			if counts > 1 {
				return start
			} else {
				return -1
			}
		}
		if counts < (middle - start + 1) {
			start = middle + 1
		} else {
			end = middle
		}
	}
	return -1
}

//
//  count
//  @Description: 统计数组中元素取值在start与end之间的元素个数
//  @param nums：
//  @param start
//  @param end
//  @return int: 元素个数
//
func count(nums []int, start, end int) int {
	total := 0
	for _, v := range nums {
		if v >= start && v <= end {
			total++
		}
	}
	return total
}

//
// Find_duplicate_v3
//  @Description: 数组中只存在一个重复元素，在o(n),o(1)的时间，空间复杂度上解决
//  @param nums
//  @return int
//
func Find_duplicate_v3(nums []int) int {
	fast, slow := nums[nums[0]], nums[0]
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return fast
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(Find_duplicate_v3(nums))
}
