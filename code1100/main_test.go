package main

import "testing"

func Test_numKLenSubstrNoRepeats(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{
			name: "case1",
			args: args{
				s: "havefunonleetcode",
				k: 5,
			},
			wantRet: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := numKLenSubstrNoRepeats(tt.args.s, tt.args.k); gotRet != tt.wantRet {
				t.Errorf("numKLenSubstrNoRepeats() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
