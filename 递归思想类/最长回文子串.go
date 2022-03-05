package 递归思想类

/*
p(i,j)=p(i+1,j-1)& si==sj
*/
func longestPalindrome(str string) string {
	strLen := len(str)
	if strLen < 2 {
		return str
	}

	dp := make([][]bool, strLen)
	for i := 0; i < strLen; i++ {
		dp[i] = make([]bool, strLen)
		dp[i][i] = true
	}

	maxLen := 1
	maxBegin := 0

	for j := 1; j < strLen; j++ {
		for i := 0; i < j; i++ {
			if str[i] != str[j] {
				dp[i][j] = false
			} else {
				if (j-i) < 3 || dp[i+1][j-1] {
					dp[i][j] = true
					if (j - i + 1) > maxLen {
						maxLen = j - i + 1
						maxBegin = i
					}
				} else {
					dp[i][j] = false
				}
			}
		}
	}

	return str[maxBegin : maxBegin+maxLen]
}
