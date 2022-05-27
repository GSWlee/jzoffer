package _6

func singleNumbers(nums []int) []int {
	z := nums[0]
	for i := 1; i < len(nums); i++ {
		z = z ^ nums[i]
	}
	t := 1
	for t&z == 0 {
		t = t << 1
	}
	nums1, nums2 := []int{}, []int{}
	for _, v := range nums {
		if v&t == 0 {
			nums1 = append(nums1, v)
		} else {
			nums2 = append(nums2, v)
		}
	}
	t1, t2 := nums1[0], nums2[0]
	for i := 1; i < len(nums1); i++ {
		t1 = t1 ^ nums1[i]
	}
	for i := 1; i < len(nums2); i++ {
		t2 = t2 ^ nums2[i]
	}
	return []int{t1, t2}
}
