/**
 * @Author: alessonhu
 * @Description:
 * @File:  字母编码成数字，多少种译码方式
 * @Version: 1.0.0
 * @Date: 2022/4/22 20:57
 */

package main

import "fmt"

// a:1 b:2 ... z:26
// in: 121 ans: 3
// in: 282 ans: 1
// 121X X=0,f(n)=f(n-2) X!=0,f(n)=f(n-1)+f(n-2)
func numsOfTrans(num string) int {
	if len(num) <= 1 {
		return len(num)
	}
	pre, cur := 1, 1
	for i := 1; i < len(num); i++ {
		tmp := cur
		if num[i] == '0' {
			if num[i-1] == '1' || num[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if num[i-1] == '1' || (num[i-1] == '2' && num[i] <= '6') {
			cur = cur + pre
		}
		pre = tmp
	}
	return cur
}

func main() {
	fmt.Println(numsOfTrans("121"))
	fmt.Println(numsOfTrans("282"))
}
