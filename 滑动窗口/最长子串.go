package 滑动窗口

import "math"

// 12334567 ans:5
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	dataMap := make(map[byte]int)
	dataMap[s[0]] = 1
	ans := 1
	for i, j := 0, 1; i < len(s) && j < len(s); {
		if _, ok := dataMap[s[j]]; ok {
			delete(dataMap, s[i])
			i++
		} else {
			dataMap[s[j]] = 1
			ans = int(math.Max(float64(ans), float64(j-i)+1))
			j++
		}
	}
	return ans
}
