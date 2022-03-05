package 递归思想类

/*
 f(n)=f(n-1)+f(n-2)
1 2 3 5 8 13

f6=f5+f4= f4+f3 + f3+f2
*/
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs1(n int) int {
	x, y := 0, 0
	ans := 1
	for i := 1; i <= n; i++ {
		x = y
		y = ans
		ans = x + y
	}
	return ans
}
