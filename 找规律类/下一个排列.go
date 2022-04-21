/**
 * @Author: alessonhu
 * @Description:
 * @File:  下一个排列
 * @Version: 1.0.0
 * @Date: 2022/4/21 22:52
 */

package 找规律类

func nextPermutation(nums []int) {
	//123465
	//123463
	//654321
	// 1 2 3
	ascI, ascJ := 0, 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			ascI = i - 1
			ascJ = i
			break
		}
	}

	secondMax := ascJ
	for i := len(nums) - 1; i >= ascJ; i-- {
		if nums[i] > nums[ascI] {
			secondMax = i
			break
		}
	}

	nums[ascI], nums[secondMax] = nums[secondMax], nums[ascI]

	start, end := ascJ, len(nums)-1
	for start <= end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

	return
}
