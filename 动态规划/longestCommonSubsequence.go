/**
 * @Author: alessonhu
 * @Description:
 * @File:  longestCommonSubsequence
 * @Version: 1.0.0
 * @Date: 2022/4/22 15:50
 */

package main

func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubstring(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	//array[i][j]中存储首字母是text[i]以及text[j]的公共子串
	array := make([][]int, len1)
	for i := 0; i < len1; i++ {
		array[i] = make([]int, len2)
	}
	//初始化边缘位置
	for i := 0; i < len1; i++ {
		if text1[i] == text2[len2-1] {
			array[i][len2-1] = 1
		} else {
			array[i][len2-1] = 0
		}
	}
	for j := 0; j < len2-1; j++ {
		if text2[j] == text1[len1-1] {
			array[len1-1][j] = 1
		} else {
			array[len1-1][j] = 0
		}
	}
	//从右下角开始
	for i := len1 - 2; i >= 0; i-- {
		for j := len2 - 2; j >= 0; j-- {
			if text1[i] == text2[j] {
				array[i][j] = array[i+1][j+1] + 1
			} else {
				array[i][j] = 0
			}
		}
	}
	max := 0
	//遍历数组，取出最大值
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if array[i][j] > max {
				max = array[i][j]
			}
		}
	}
	return max
}
