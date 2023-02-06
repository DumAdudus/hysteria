package obfs

import (
	"github.com/valyala/bytebufferpool"
)

const FlipTrigger = "BitsFlip"

var _ Obfuscator = &BitFlipObfuscator{}

type BitFlipObfuscator struct{}

func NewFlipObfuscator() *BitFlipObfuscator {
	return &BitFlipObfuscator{}
}

func flipBits(in []byte, out []byte) int {
	for i, c := range in {
		out[i] = ^c
	}
	return len(in)
}

func (x *BitFlipObfuscator) Deobfuscate(in []byte, out []byte) int {
	return flipBits(in, out)
}

func (x *BitFlipObfuscator) Obfuscate(in []byte, out []byte) int {
	return flipBits(in, out)
}

func (x *BitFlipObfuscator) ObfuscateOnBuffer(in []byte, out *bytebufferpool.ByteBuffer) int {
	for _, c := range in {
		out.WriteByte(^c)
	}
	return len(in)
}
