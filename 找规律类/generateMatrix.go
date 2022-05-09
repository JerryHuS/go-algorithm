/**
 * @Author: alessonhu
 * @Description: 生成扫雷地图，富途三面算法题
 * @File:  generateMatrix
 * @Version: 1.0.0
 * @Date: 2022/4/27 16:59
 */
package 找规律类

import (
	"fmt"
	"math/rand"
)

func main() {
	res := generateMatrix(5, 5, 5)
	for i := range res {
		fmt.Println(res[i])
	}
}

func generateMatrix(m, n, k int) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for k > 0 {
		posI := rand.Intn(m)
		posJ := rand.Intn(n)
		if matrix[posI][posJ] == -1 {
			continue
		}
		matrix[posI][posJ] = -1
		k--
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != -1 {
				matrix[i][j] = calPosNum(i, j, matrix)
			}
		}
	}
	return matrix
}

func calPosNum(i, j int, matrix [][]int) int {
	ans := 0
	for m := i - 1; m <= i+1; m++ {
		for n := j - 1; n <= j+1; n++ {
			if m >= 0 && m < len(matrix) && n >= 0 && n < len(matrix[0]) && matrix[m][n] == -1 {
				ans++
			}
		}
	}
	return ans
}
