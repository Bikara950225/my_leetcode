package code_mianshi

import "testing"

func TestPrintMsg(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintMsg(tt.args.msg)
		})
	}
}
