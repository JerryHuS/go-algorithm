/**
 * @Author: alessonhu
 * @Description:
 * @File:  字典序排数
 * @Version: 1.0.0
 * @Date: 2022/4/18 16:25
 */

package 找规律类

func lexicalOrder(n int) []int {
	arr := make([]int, n)
	num := 1
	for i := range arr {
		arr[i] = num
		if num*10 < n {
			num = num * 10
		} else {
			for num%10 == 9 || num+1 > n {
				num = num / 10
			}
			num++
		}
	}
	return arr
}
