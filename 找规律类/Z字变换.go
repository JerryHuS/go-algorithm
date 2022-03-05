package 找规律类

import "bytes"

/*0 1 2 3 4 5 6 7 8
0 P     I     N
1 A   L S   I G
2 Y A   H R
3 P     I

len
n

周期：n+n-2
列数：1+n-2

len/2n-2 * n-1
*/

func convert1(s string, numRows int) string {
	dp := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]byte, len(s))
	}

	sIndex := 0
	for j := 0; j < len(s); j++ {
		for i := 0; i < numRows; i++ {
			if (j%(numRows-1) == 0 || (i+j)%(numRows-1) == 0) && sIndex < len(s) {
				dp[i][j] = s[sIndex]
				sIndex++
			} else {
				dp[i][j] = 0
			}
		}
	}

	result := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			if dp[i][j] != 0 {
				result = append(result, dp[i][j])
			}
		}
	}

	return string(result)
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	mat := make([][]byte, numRows)
	circle := numRows + numRows - 2
	x := 0
	for i, _ := range s {
		mat[x] = append(mat[x], s[i])
		if i%circle < numRows-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}
