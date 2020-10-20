package util

import (
	"testing"
)

func TestEncodeMd5(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"zKen",
			args{value: "go_qrcode_login"},
			"c465801b2e6c733a6eb9885f2a37d840",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeMd5(tt.args.value); got != tt.want {
				t.Errorf("EncodeMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodePasswd(t *testing.T) {
	EncodePasswd([]byte("shiluonb"))
}
