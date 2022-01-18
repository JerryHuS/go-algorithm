package 优先搜索

/*
1 1 1
1 1 0
1 0 1

1,1,2

2 2 2
2 2 0
2 0 1
*/

var (
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}
)

//BFS
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	n, m := len(image), len(image[0])
	queue := [][]int{{sr, sc}}
	curColor := image[sr][sc]
	image[sr][sc] = newColor
	for i := 0; i < len(queue); i++ {
		for j := 0; j < 4; j++ {
			mx, my := queue[i][0]+dx[j], queue[i][1]+dy[j]
			if mx >= 0 && mx < n && my >= 0 && my < m && image[mx][my] == curColor {
				image[mx][my] = newColor
				queue = append(queue, []int{mx, my})
			}
		}
	}
	return image
}

//DFS 回溯
func floodFillDFS(image [][]int, sr int, sc int, newColor int) [][]int {
	currColor := image[sr][sc]
	if currColor != newColor {
		floodDFS(image, sr, sc, currColor, newColor)
	}
	return image
}

func floodDFS(image [][]int, x, y, color, newColor int) {
	if image[x][y] == color {
		image[x][y] = newColor
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx >= 0 && my >= 0 && mx < len(image) && my < len(image[0]) {
				floodDFS(image, mx, my, color, newColor)
			}
		}
	}
}
