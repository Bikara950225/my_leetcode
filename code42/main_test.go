package main

import "testing"

func Test_trapWithDP(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{
			args: args{
				height: []int{
					0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1,
				},
			},
			wantRet: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := trapWithDP(tt.args.height); gotRet != tt.wantRet {
				t.Errorf("trapWithDP() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_trapWithStack(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{
			args: args{
				height: []int{
					0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1,
				},
			},
			wantRet: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := trapWithStack(tt.args.height); gotRet != tt.wantRet {
				t.Errorf("trapWithStack() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
