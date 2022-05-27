package _6

func singleNumber(nums []int) int {
	bits := [64]int{}
	for i := 0; i < len(nums); i++ {
		t := 1
		for j := 0; j < 64; j++ {
			if t&nums[i] != 0 {
				bits[j]++
			}
			t = t << 1
		}
	}
	ans := 0
	for i := 63; i >= 0; i-- {
		ans = ans << 1
		ans += bits[i] % 3
	}
	return ans
}
