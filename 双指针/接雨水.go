package 双指针

import "math"

//暴力法，算每个节点 左右的最大值，然后算最大接水量
func trap(height []int) int {
	ans := 0
	for i := 1; i < len(height); i++ {
		maxLeft, maxRight := 0, 0
		for left := i; left >= 0; left-- {
			maxLeft = int(math.Max(float64(height[left]), float64(maxLeft)))
		}
		for right := i; right < len(height); right++ {
			maxRight = int(math.Max(float64(height[right]), float64(maxRight)))
		}
		ans += int(math.Min(float64(maxLeft), float64(maxRight))) - height[i]
	}
	return ans
}

/*
双指针
4, 2, 0, 3, 2, 5
 			  _
 _           | |
| |     _    | |
| | _  | | _ | |
| || | | || || |
|_||_|_|_||_||_|
 i            j
*/
func trap1(height []int) int {
	ans := 0
	begin, end := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for begin < end {
		if height[begin] < height[end] {
			if height[begin] > leftMax {
				leftMax = height[begin]
			} else {
				ans += leftMax - height[begin]
			}
			begin++
		} else {
			if height[end] > rightMax {
				rightMax = height[end]
			} else {
				ans += rightMax - height[end]
			}
			end--
		}
	}
	return ans
}
