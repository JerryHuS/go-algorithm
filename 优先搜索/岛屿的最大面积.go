package 优先搜索

import (
	"container/list"
	"math"
)

/*
输入：grid = [
			[0,0,1,0,0,0,0,1,0,0,0,0,0],
			[0,0,0,0,0,0,0,1,1,1,0,0,0],
			[0,1,1,0,1,0,0,0,0,0,0,0,0],
			[0,1,0,0,1,1,0,0,1,0,1,0,0],
			[0,1,0,0,1,1,0,0,1,1,1,0,0],
			[0,0,0,0,0,0,0,0,0,0,1,0,0],
			[0,0,0,0,0,0,0,1,1,1,0,0,0],
			[0,0,0,0,0,0,0,1,1,0,0,0,0]
			]
输出：6
解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1 。
*/

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ans = int(math.Max(float64(ans), float64(maxDFS(grid, i, j))))
		}
	}
	return ans
}

func maxDFS(grid [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != 1 {
		return 0
	}
	grid[x][y] = 0
	ans := 1
	for j := 0; j < 4; j++ {
		mx, my := x+dx[j], y+dy[j]
		ans += maxDFS(grid, mx, my)
	}
	return ans
}

func maxAreaOfIslandBFS(grid [][]int) int {
	ans := 0
	n, m := len(grid), len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			cur := 0
			queue := list.New()
			queue.PushBack([2]int{i, j})
			for queue.Len() > 0 {
				qNode := queue.Remove(queue.Front()).([2]int)
				x, y := qNode[0], qNode[1]
				if x < 0 || y < 0 || x >= n || y >= m || grid[x][y] != 1 {
					continue
				}
				grid[x][y] = 0
				cur++
				for k := 0; k < 4; k++ {
					queue.PushBack([2]int{x + dx[k], y + dy[k]})
				}
			}
			ans = int(math.Max(float64(ans), float64(cur)))
		}
	}
	return ans
}
