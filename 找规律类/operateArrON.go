/**
 * @Author: alessonhu
 * @Description: 主要是字节三面遇到的一些题目，考察时间或空间限制下下对数组的一些操作
 * @File:  operateArrON
 * @Version: 1.0.0
 * @Date: 2022/4/25 10:54
 */
package 找规律类

import (
	"container/list"
	"math"
	"sort"
)

/*
找出数组中的元素，满足该元素左边的都比该元素小，右边的都比该元素大
{1,3,2,4,7,9,8} 4 7
https://juejin.cn/post/6941648141049397261
*/
func findElemBetween(input []int) []int {
	if len(input) <= 2 {
		return []int{}
	}
	n := len(input)

	minRight := make([]int, n)

	for i := n - 2; i >= 0; i-- {
		if i+2 <= n-1 {
			minRight[i] = minInt(input[i+1], minRight[i+1])
		} else {
			minRight[i] = input[i+1]
		}
	}

	maxLeft := input[0]
	var res []int
	for i := 1; i < n; i++ {
		if input[i] > maxLeft {
			maxLeft = input[i]
			if input[i] < minRight[i] {
				res = append(res, input[i])
			}
		}
	}

	return res
}

// 带条件的单调栈
// 1 4 7
func findElemBetween1(input []int) []int {
	queue := list.New()
	max := math.MinInt32
	for i := 0; i < len(input); i++ {
		for queue.Len() > 0 && input[i] <= queue.Back().Value.(int) {
			queue.Remove(queue.Back())
		}
		if input[i] > max {
			queue.PushBack(input[i])
		}
		max = maxInt(max, input[i])
	}

	var res []int
	for queue.Len() > 0 {
		res = append(res, queue.Remove(queue.Back()).(int))
	}

	return res
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Input:
ar1[] = {1, 5, 9, 10, 15, 20};
ar2[] = {2, 3, 8, 13};

Output:
ar1[] = {1, 2, 3, 5, 8, 9}
ar2[] = {10, 13, 15, 20}

O(1)空间
*/

func sortTwoArr(arr1 []int, arr2 []int) {
	arr1Len := len(arr1)
	arr2Len := len(arr2)

	i, j, k := 0, 0, arr1Len-1
	for i <= k && j < arr2Len {
		if arr1[i] < arr2[j] {
			i++
		} else {
			arr2[j], arr1[k] = arr1[k], arr2[j]
			j++
			k--
		}
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	return
}
