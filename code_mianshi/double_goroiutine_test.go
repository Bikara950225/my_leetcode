package code_mianshi

import "testing"

func Test_doubleG(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "happy path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//doubleG()
			//doubleG2()
			doubleGT()
		})
	}
}
