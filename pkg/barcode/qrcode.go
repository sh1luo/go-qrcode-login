package barcode

import (
	"bytes"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/jpeg"
)

type QrCode struct {
	Content string
	Height  int
	Width   int
	Level   qr.ErrorCorrectionLevel
	Mode    qr.Encoding
}

// ps:Trying to implement the Writer interface for the []byte type
// so that it can be written by jpeg.Encode, suddenly remembered that there is already :)
//
//type bytesContainer []byte
//
//func (b *bytesContainer) Write(p []byte) (n int, err error) {
//
//}

func (q *QrCode) Encode() ([]byte, error) {
	bc, err := qr.Encode(q.Content, q.Level, q.Mode)
	if err != nil {
		return nil, err
	}
	bc, err = barcode.Scale(bc, q.Width, q.Height)

	// var retBytes []byte
	var buffer bytes.Buffer
	err = jpeg.Encode(&buffer, bc, nil)

	if err != nil || buffer.Len() == 0 {
		return nil, err
	}

	return buffer.Bytes(), nil
}
