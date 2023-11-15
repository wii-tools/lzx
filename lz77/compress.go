package lz77

import "github.com/wii-tools/lzx/lz10"

func Compress(data []byte) ([]byte, error) {
	compressed, err := lz10.Compress(data)
	if err != nil {
		return nil, err
	}

	return append(LZ77Magic[:], compressed...), nil
}
