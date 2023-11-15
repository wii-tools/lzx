package lz77

import "errors"

var (
	LZ77Magic = [4]byte{'L', 'Z', '7', '7'}

	InvalidMagicError = errors.New("invalid file magic")
)
