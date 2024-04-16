package code_mianshi

import "testing"

func Test_strIpV4ToNumber(t *testing.T) {
	type args struct {
		ipv4 string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ipv4: "172.30.5.35",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := strIpV4ToNumber(tt.args.ipv4)
			if (err != nil) != tt.wantErr {
				t.Errorf("strIpV4ToNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotIpStr := uint32ToStrIp(got)
			if gotIpStr != tt.args.ipv4 {
				t.Errorf("strIpV4ToNumber() error: ipv4 != gotIpStr: %s != %s", gotIpStr, tt.args.ipv4)
				return
			}
			gotIpStr2 := uint32ToStrIp2(got)
			if gotIpStr2 != tt.args.ipv4 {
				t.Errorf("strIpV4ToNumber() error: ipv4 != gotIpStr2: %s != %s", gotIpStr2, tt.args.ipv4)
				return
			}
		})
	}
}
