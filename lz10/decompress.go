package lz10

import (
	"bytes"
	"encoding/binary"
)

func Decompress(data []byte) ([]byte, error) {
	compressed := bytes.NewBuffer(data)

	// Ensure size validity.
	if CompressedMin > compressed.Len() {
		return nil, ErrCompressedTooSmall
	}
	if compressed.Len() > CompressedMax {
		return nil, ErrCompressedTooLarge
	}
	// Ensure the first byte of this data is 0x11, signifying a proper file.
	if compressed.Next(1)[0] != FileMagic {
		return nil, ErrInvalidMagic
	}

	// Obtain the length of the decompressed file.
	// We then drop the highest byte to strip the 0x10 magic.
	header := append(compressed.Next(3), []byte{0}...)
	originalLen := binary.LittleEndian.Uint32(header)

	decompressed := new(bytes.Buffer)
	var mask byte
	var flags byte
	var pos int
	var length int

	for decompressed.Len() < int(originalLen) {
		mask >>= BitShiftCount
		if mask == 0 {
			flags = compressed.Next(1)[0]
			mask = DefaultMask
		}

		if (flags & mask) == 0 {
			decompressed.WriteByte(compressed.Next(1)[0])
		} else {
			pos = int(compressed.Next(1)[0])
			pos = (pos << 8) | int(compressed.Next(1)[0])
			length = (pos >> 12) + Threshold + 1
			pos = (pos & 0xFFF) + 1

			for length != 0 {
				decompressed.WriteByte(decompressed.Bytes()[decompressed.Len()-pos])
				length--
			}
		}
	}

	return decompressed.Bytes(), nil
}
