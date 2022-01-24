package 优先搜索

import (
	"container/list"
)

// 0 1矩阵 到0最近的距离
func updateMatrix(mat [][]int) [][]int {
	n, m := len(mat), len(mat[0])
	queue := list.New()
	meetMap := make(map[[2]int]bool)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				queue.PushBack([2]int{i, j})
				meetMap[[2]int{i, j}] = true
			}
		}
	}

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).([2]int)
		for j := 0; j < 4; j++ {
			mx, my := node[0]+dx[j], node[1]+dy[j]
			if mx >= 0 && my >= 0 && mx < n && my < m && !meetMap[[2]int{mx, my}] {
				mat[mx][my] = mat[node[0]][node[1]] + 1
				meetMap[[2]int{mx, my}] = true
				queue.PushBack([2]int{mx, my})
			}
		}
	}
	return mat
}
