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
		{
			name: "happy path",
			args: args{
				msg: "hello",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintMsg(tt.args.msg)
		})
	}
}
