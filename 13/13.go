package _3

func sum(a, b int) int {
	s := 0
	for a != 0 {
		s = s + a%10
		a = a / 10
	}
	for b != 0 {
		s = s + b%10
		b = b / 10
	}
	return s
}

type loc struct {
	x int
	y int
}

func movingCount(m int, n int, k int) int {
	flag := [][]int{}
	for i := 0; i < m; i++ {
		flag = append(flag, make([]int, n))
	}
	ans := 0
	dp := []loc{}
	if sum(0, 0) <= k {
		ans++
		flag[0][0] = 1
		dp = append(dp, loc{
			x: 0,
			y: 0,
		})
	}
	for len(dp) > 0 {
		tx, ty := dp[0].x, dp[0].y
		dp = dp[1:]
		if tx+1 < m && flag[tx+1][ty] == 0 {
			if sum(tx+1, ty) <= k {
				flag[tx+1][ty] = 1
				ans++
				dp = append(dp, loc{
					x: tx + 1,
					y: ty,
				})
			} else {
				flag[tx+1][ty] = 2
			}
		}
		if tx-1 > 0 && flag[tx-1][ty] == 0 {
			if sum(tx-1, ty) <= k {
				flag[tx-1][ty] = 1
				ans++
				dp = append(dp, loc{
					x: tx - 1,
					y: ty,
				})
			} else {
				flag[tx-1][ty] = 2
			}
		}
		if ty+1 < n && flag[tx][ty+1] == 0 {
			if sum(tx, ty+1) <= k {
				flag[tx][ty+1] = 1
				ans++
				dp = append(dp, loc{
					x: tx,
					y: ty + 1,
				})
			} else {
				flag[tx][ty+1] = 2
			}
		}
		if ty-1 > 0 && flag[tx][ty-1] == 0 {
			if sum(tx, ty-1) <= k {
				flag[tx][ty-1] = 1
				ans++
				dp = append(dp, loc{
					x: tx,
					y: ty - 1,
				})
			} else {
				flag[tx][ty-1] = 2
			}
		}
	}
	return ans
}
