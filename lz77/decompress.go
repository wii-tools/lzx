package lz77

import (
	"bytes"
	"github.com/wii-tools/lzx/lz10"
)

func Decompress(passed []byte) ([]byte, error) {
	if bytes.Equal(passed[:4], LZ77Magic[:]) {
		return nil, InvalidMagicError
	}

	return lz10.Decompress(passed)
}
