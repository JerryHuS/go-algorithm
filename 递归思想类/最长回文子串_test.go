package 递归思想类

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{str: "eabcba"}, want: "abcba"},
		{name: "case2", args: args{str: "a"}, want: "a"},
		{name: "case3", args: args{str: "aa"}, want: "aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.str); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
