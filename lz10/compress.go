package lz10

import (
	"bytes"
	"encoding/binary"
	"math"
)

type compressionContext struct {
	decompressed *[]byte
	position     int
	index        int
	length       int
}

func (c *compressionContext) Search(length, position *int) {
	*length = Threshold

	if c.index >= MaxOffset {
		c.position = MaxOffset
	} else {
		c.position = c.index
	}

	for ; c.position > VRAMCompatible; c.position-- {
		for c.length = 0; c.length < MaxCoded; c.length++ {
			if c.index+c.length == len(*c.decompressed) {
				break
			}

			if c.index+c.length >= len(*c.decompressed) {
				break
			}

			if (*c.decompressed)[c.index+c.length] != (*c.decompressed)[c.index+c.length-c.position] {
				break
			}
		}

		if c.length > *length {
			*position = c.position
			*length = c.length

			if *length == MaxCoded {
				break
			}
		}
	}
}

func Compress(data []byte) ([]byte, error) {
	compressed := new(bytes.Buffer)
	err := binary.Write(compressed, binary.LittleEndian, uint32(FileMagic|len(data)<<8))
	if err != nil {
		return nil, ErrFailedBufferWrite
	}

	var ctx compressionContext
	ctx.decompressed = &data

	var mask int
	var flag int
	var lenBest int
	var posBest int
	var lenNext int
	var posNext int
	var lenPost int
	var posPost int

	for ctx.index < len(data) {
		mask >>= BitShiftCount
		if mask == 0 {
			compressed.WriteByte(0)

			flag = compressed.Len() - 1
			mask = DefaultMask
		}

		ctx.Search(&lenBest, &posBest)

		if ctx.index+lenBest < len(data) {
			ctx.index += lenBest
			ctx.Search(&lenNext, &posNext)
			ctx.index -= lenBest - 1
			ctx.Search(&lenPost, &posPost)
			ctx.index--

			if lenNext <= Threshold {
				lenNext = 1
			}

			if lenPost <= Threshold {
				lenPost = 1
			}

			if lenBest+lenNext <= 1+lenPost {
				lenBest = 1
			}
		}

		if lenBest > Threshold {
			ctx.index += lenBest
			compressed.Bytes()[flag] |= byte(mask)
			_bytes := []byte{
				byte(((lenBest - (Threshold + 1)) << 4) | ((posBest - 1) >> 8)),
				byte((posBest - 1) & math.MaxUint8),
			}

			compressed.Write(_bytes)
		} else {
			compressed.WriteByte(data[ctx.index])
			ctx.index++
		}
	}

	return compressed.Bytes(), nil
}
