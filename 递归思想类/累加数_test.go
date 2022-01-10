/**
 * @Author: alessonhu
 * @Description:
 * @File:  累加数_test.go
 * @Version: 1.0.0
 * @Date: 2022/1/10 下午9:40
 */

package 递归思想类

import "testing"

func Test_isAdditiveNumber(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{num: "123"}, want: true},
		{name: "case2", args: args{num: "124"}, want: false},
		{name: "case3", args: args{num: "112358"}, want: true},
		{name: "case4", args: args{num: "199100199"}, want: true},
		{name: "case5", args: args{num: "000"}, want: true},
		{name: "case6", args: args{num: "111"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdditiveNumber(tt.args.num); got != tt.want {
				t.Errorf("isAdditiveNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
