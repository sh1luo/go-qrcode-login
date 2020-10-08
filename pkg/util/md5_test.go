package util

import (
	"log"
	"testing"
)

func TestEncodeMd5(t *testing.T) {
	want := "c465801b2e6c733a6eb9885f2a37d840"
	got := EncodeMd5("go_qrcode_login")
	if want != got {
		log.Printf("Md5 got!=want")
	}
}
