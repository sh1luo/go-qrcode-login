package gredis

import "encoding/binary"

//Chain
type chain struct {
	data []byte
	err  error
}

func (b *chain) MustUint32() uint32 {
	return binary.BigEndian.Uint32(b.data)
}

func (b *chain) MustString() string {
	return string(b.data)
}

func (b *chain) Error() string {
	return b.err.Error()
}
