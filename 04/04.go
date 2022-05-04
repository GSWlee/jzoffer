package main

//
// Find_number
//  @Description: 判断数组中是否含有目标元素
//  @param arr
//  @param target
//  @return bool
//
func Find_number(arr [][]int, target int) bool {
	if len(arr) == 0 {
		return false
	}
	if len(arr[0]) == 0 {
		return false
	}
	h, w := len(arr), len(arr[0])
	x, y := 0, w-1
	for x < h && y >= 0 {
		if arr[x][y] == target {
			return true
		} else if arr[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
