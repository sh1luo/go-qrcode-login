package util

import (
	"testing"
	"time"
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
	//var i = 0
	//var dst []byte
	//for i<10 {
	//
	//	dst = append(dst,[]byte(src)[:8]...)
	//	fmt.Println()
	//	i++
	//}
}

func BenchmarkGetRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Base64EncodeToString([]byte(time.Now().String()))
	}
}

func BenchmarkGetRandomString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetRandomString(9)
	}
}
