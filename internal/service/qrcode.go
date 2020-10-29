package service

import (
	"github.com/boombuler/barcode/qr"
	"github.com/sh1luo/go-qrcode-login.git/global"
	"github.com/sh1luo/go-qrcode-login.git/pkg/barcode"
	"github.com/sh1luo/go-qrcode-login.git/pkg/errcode"
)

func (svc *Service) GetQrCode(aftEncode string) (qrcode []byte, err *errcode.Error) {
	//st := strconv.FormatInt(time.Now().UnixNano(), 10)
	//var src []byte
	//src = append(src, []byte(param)...)
	//src = append(src, []byte(st)...)
	//aftEncode := util.Base64EncodeToString(src)
	//aftEncode = aftEncode[:10] + aftEncode[len(aftEncode)-2:]

	send := global.ServerSetting.Addr + global.ServerSetting.HttpPort + "/l/" + aftEncode
	qrc := barcode.QrCode{
		Content: send,
		Height:  200,
		Width:   200,
		Level:   qr.M,
		Mode:    qr.Auto,
	}
	var e error
	qrcode, e = qrc.Encode()
	if e != nil {
		err = errcode.ServerInternalErr
	}

	return qrcode, err
}
