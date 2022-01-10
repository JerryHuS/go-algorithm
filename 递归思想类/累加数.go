/**
 * @Author: alessonhu
 * @Description:
示例 1：

输入："112358"
输出：true
解释：累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
示例 2：

输入："199100199"
输出：true
解释：累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199


提示：

1 <= num.length <= 35
num 仅由数字（0 - 9）组成

 * @File:  func.go
 * @Version: 1.0.0
 * @Date: 2022/1/10 下午9:27
*/

package 递归思想类

import (
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	if len(num) == 0 {
		return false
	}
	lenNum := len(num)
	for i := 1; i <= lenNum/2; i++ {
		if num[0] == '0' && i > 1 {
			return false
		}
		n1, _ := strconv.Atoi(num[:i])
		for j := 1; i+j < lenNum; j++ {
			if num[i] == '0' && j > 1 {
				break
			}
			n2, _ := strconv.Atoi(num[i : i+j])
			sufNum := num[i+j:]
			if isMatch(n1, n2, sufNum) {
				return true
			}
		}
	}
	return false
}

func isMatch(n1, n2 int, num string) bool {
	if len(num) == 0 {
		return true
	}
	sumStr := strconv.Itoa(n1 + n2)
	if !strings.HasPrefix(num, sumStr) {
		return false
	}
	return isMatch(n2, n1+n2, num[len(sumStr):])
}
