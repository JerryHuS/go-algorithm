package 滑动窗口

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		//{name: "case1", args: args{s1: "acb", s2: "babce"}, want: true},
		//{name: "case2", args: args{s1: "abc", s2: "abedc"}, want: false},
		{name: "case3", args: args{s1: "adc", s2: "dcda"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
