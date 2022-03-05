package main

import "testing"

func Test_getIpType(t *testing.T) {
	type args struct {
		addrIp string
	}
	tests := []struct {
		name          string
		args          args
		wantIpType    string
		wantIsPrivate bool
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{addrIp: "1.0.0.1"}, wantIpType: "A", wantIsPrivate: false},
		{name: "case2", args: args{addrIp: "192.168.0.2"}, wantIpType: "C", wantIsPrivate: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIpType, gotIsPrivate := getIpType(tt.args.addrIp)
			if gotIpType != tt.wantIpType {
				t.Errorf("getIpType() gotIpType = %v, want %v", gotIpType, tt.wantIpType)
			}
			if gotIsPrivate != tt.wantIsPrivate {
				t.Errorf("getIpType() gotIsPrivate = %v, want %v", gotIsPrivate, tt.wantIsPrivate)
			}
		})
	}
}

func Test_isMaskValid(t *testing.T) {
	type args struct {
		addrMask string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "c1", args: args{"255.16.0.0"}, want: false},
		//{name: "c2", args: args{"255.255.255.0"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMaskValid(tt.args.addrMask); got != tt.want {
				t.Errorf("isMaskValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
