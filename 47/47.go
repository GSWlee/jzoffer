package _7

func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 {
				if j != 0 {
					grid[i][j] = grid[i][j] + grid[i][j-1]
				}
			} else {
				if j == 0 {
					grid[i][j] = grid[i-1][j] + grid[i][j]
				} else {
					if grid[i][j-1] > grid[i-1][j] {
						grid[i][j] = grid[i][j-1] + grid[i][j]
					} else {
						grid[i][j] = grid[i][j] + grid[i-1][j]
					}
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
