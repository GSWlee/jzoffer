package _2

import "fmt"

func find(board [][]byte, flag [][]bool, word string, now, x, y int) bool {
	if len(word) == now {
		return true
	}
	target := word[now]
	if x == -1 {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] == target {
					flag[i][j] = true
					f := find(board, flag, word, now+1, i, j)
					if f == true {
						return true
					} else {
						flag[i][j] = false
					}
				}
			}
		}
	} else {
		xmax, ymax := len(board), len(board[0])
		if x-1 >= 0 && flag[x-1][y] == false && board[x-1][y] == target {
			flag[x-1][y] = true
			f := find(board, flag, word, now+1, x-1, y)
			if f == true {
				return true
			} else {
				flag[x-1][y] = false
			}
		}
		if y+1 < ymax && flag[x][y+1] == false && board[x][y+1] == target {
			flag[x][y+1] = true
			f := find(board, flag, word, now+1, x, y+1)
			if f == true {
				return true
			} else {
				flag[x][y+1] = false
			}
		}
		if x+1 < xmax && flag[x+1][y] == false && board[x+1][y] == target {
			flag[x+1][y] = true
			f := find(board, flag, word, now+1, x+1, y)
			if f == true {
				return true
			} else {
				flag[x+1][y] = false
			}
		}
		if y-1 >= 0 && flag[x][y-1] == false && board[x][y-1] == target {
			flag[x][y-1] = true
			f := find(board, flag, word, now+1, x, y-1)
			if f == true {
				return true
			} else {
				flag[x][y-1] = false
			}
		}
		return false

	}
	return false
}

func exist(board [][]byte, word string) bool {
	flag := [][]bool{}
	for _, v := range board {
		flag = append(flag, make([]bool, len(v)))
	}
	return find(board, flag, word, 0, -1, -1)
}

func main() {
	t := [][]byte{}
	t = append(t, []byte{'c', 'a', 'a'})
	t = append(t, []byte{'a', 'a', 'a'})
	t = append(t, []byte{'b', 'c', 'd'})
	a := exist(t, "aab")
	fmt.Println(a)
}
