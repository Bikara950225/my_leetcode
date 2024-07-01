package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			wantRet: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := maxProfit(tt.args.prices); gotRet != tt.wantRet {
				t.Errorf("maxProfit() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
