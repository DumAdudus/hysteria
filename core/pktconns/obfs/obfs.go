package obfs

import (
	"encoding/binary"
	"math/rand"
	"sync/atomic"

	"github.com/valyala/bytebufferpool"
	hasher "github.com/zeebo/blake3"
)

type Obfuscator interface {
	Deobfuscate(in []byte, out []byte) int
	Obfuscate(in []byte, out []byte) int
}

const (
	xpSaltLen  = 4 // bytes of uint32
	xorKeySize = 32
)

var _ Obfuscator = &XPlusObfuscator{}

// XPlusObfuscator obfuscates payload using one-time keys generated from hashing a pre-shared key and random salt.
// Packet format: [salt][obfuscated payload]
type XPlusObfuscator struct {
	Key  []byte
	salt atomic.Uint32
}

func NewXPlusObfuscator(key []byte) *XPlusObfuscator {
	obfs := &XPlusObfuscator{
		Key:  key,
		salt: atomic.Uint32{},
	}
	obfs.salt.Store(rand.Uint32() & 0xFFF)
	return obfs
}

func (x *XPlusObfuscator) Deobfuscate(in []byte, out []byte) int {
	outLen := len(in) - xpSaltLen
	if outLen <= 0 || len(out) < outLen {
		return 0
	}
	key := hasher.Sum256(append(x.Key, in[:xpSaltLen]...))
	for i, c := range in[xpSaltLen:] {
		out[i] = c ^ key[i%xorKeySize]
	}
	return outLen
}

func (x *XPlusObfuscator) Obfuscate(in []byte, out []byte) int {
	outLen := len(in) + xpSaltLen
	if len(out) < outLen {
		return 0
	}

	salt := out[:xpSaltLen]
	binary.LittleEndian.PutUint32(salt, x.salt.Add(1))
	xorKey := hasher.Sum256(append(x.Key, salt...))

	payload := out[xpSaltLen:]
	for i, c := range in {
		payload[i] = c ^ xorKey[i%xorKeySize]
	}
	return outLen
}

func (x *XPlusObfuscator) ObfuscateOnBuffer(in []byte, out *bytebufferpool.ByteBuffer) int {
	outLen := len(in) + xpSaltLen

	salt := x.salt.Add(1)
	key_salt := bytebufferpool.Get()
	defer bytebufferpool.Put(key_salt)

	key_salt.Set(x.Key)
	// little endian
	func(buf *bytebufferpool.ByteBuffer, v uint32) {
		buf.WriteByte(byte(v))
		buf.WriteByte(byte(v >> 8))
		buf.WriteByte(byte(v >> 16))
		buf.WriteByte(byte(v >> 24))
	}(key_salt, salt)
	out.Write(key_salt.Bytes()[len(x.Key):])

	xorKey := hasher.Sum256(key_salt.Bytes())
	for i, c := range in {
		out.WriteByte(c ^ xorKey[i%xorKeySize])
	}
	return outLen
}
