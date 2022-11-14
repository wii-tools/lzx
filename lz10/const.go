package lz10

import "errors"

const (
	// CompressedMin is the minimum file size a compressed file can have - only its header.
	CompressedMin = 0x00000004
	// CompressedMax is the maximum possible file size a LZ11 compressed file can have.
	// As described within lzss's source:
	// 0x01200003, padded to 20MB:
	// * header, 4
	// * length, RAW_MAXIM
	// * flags, (RAW_MAXIM + 7) / 8
	// 4 + 0x00FFFFFF + 0x00200000 + padding
	CompressedMax = 0x01400000

	// FileMagic represents the byte that should be present within all LZ10 compressed files.
	FileMagic = 0x10
	MaxOffset = 1 << 12

	// Threshold represents the max bytes not to encode
	Threshold = 2

	MaxCoded       = (1 << 4) + Threshold
	BitShiftCount  = 1
	DefaultMask    = 0x80
	VRAMCompatible = 1
)

var (
	ErrFailedBufferWrite  = errors.New("unable to write to byte buffer")
	ErrCompressedTooSmall = errors.New("passed data does not meet minimum required data size")
	ErrCompressedTooLarge = errors.New("passed data exceeds maximum possible data size")
	ErrInvalidMagic       = errors.New("passed data does not appear to be valid LZ10 data")
)
