package wechat

import (
	"encoding/binary"
	"math/rand"
	"net"
	"os"
	"syscall"
	"time"

	"github.com/apernet/hysteria/core/pktconns/obfs"
	"github.com/valyala/bytebufferpool"
)

const (
	udpBufferSize = 2048
	headerSize    = 13
)

var (
	initHeader                = [headerSize]byte{0xa1, 0x08, 0xff, 0xff, 0xff, 0xff, 0x00, 0x10, 0x11, 0x18, 0x30, 0x22, 0x30}
	hintBuf                   = make([]byte, udpBufferSize)
	_          net.PacketConn = &ObfsWeChatUDPPacketConn{}
)

// ObfsWeChatUDPPacketConn is still a UDP packet conn, but it adds WeChat video call header to each packet.
// Obfs in this case can be nil
type ObfsWeChatUDPPacketConn struct {
	orig *net.UDPConn
	obfs obfs.Obfuscator
	sn   uint32
}

func NewObfsWeChatUDPConn(orig *net.UDPConn, obfs obfs.Obfuscator) *ObfsWeChatUDPPacketConn {
	conn := &ObfsWeChatUDPPacketConn{
		orig: orig,
		obfs: obfs,
		sn:   rand.Uint32() & 0xFFF,
	}
	return conn
}

func (c *ObfsWeChatUDPPacketConn) ReadFrom(p []byte) (int, net.Addr, error) {
	poolBuf := bytebufferpool.Get()
	defer bytebufferpool.Put(poolBuf)

	for {
		poolBuf.Set(hintBuf)
		n, addr, err := c.orig.ReadFrom(poolBuf.Bytes())
		if n <= headerSize {
			return 0, addr, err
		}
		payload := poolBuf.Bytes()[headerSize:n]
		var newN int
		if c.obfs != nil {
			newN = c.obfs.Deobfuscate(payload, p)
		} else {
			newN = copy(p, payload)
		}
		if newN > 0 {
			// Valid packet
			return newN, addr, err
		} else if err != nil {
			// Not valid and orig.ReadFrom had some error
			return 0, addr, err
		}
	}
}

func (c *ObfsWeChatUDPPacketConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	header := initHeader[:]
	// Set SN, we don't care what is it, so no lock
	sn := header[2:6]
	binary.BigEndian.PutUint32(sn, c.sn)
	c.sn++

	poolBuf := bytebufferpool.Get()
	defer bytebufferpool.Put(poolBuf)
	poolBuf.Set(header)

	if c.obfs != nil {
		xplusObfs, _ := c.obfs.(*obfs.XPlusObfuscator) // Currently there's only this XPlus obfs
		xplusObfs.ObfuscateOnBuffer(p, poolBuf)
	} else {
		poolBuf.Write(p)
	}
	_, err = c.orig.WriteTo(poolBuf.Bytes(), addr)

	if err != nil {
		return 0, err
	} else {
		return len(p), nil
	}
}

func (c *ObfsWeChatUDPPacketConn) Close() error {
	return c.orig.Close()
}

func (c *ObfsWeChatUDPPacketConn) LocalAddr() net.Addr {
	return c.orig.LocalAddr()
}

func (c *ObfsWeChatUDPPacketConn) SetDeadline(t time.Time) error {
	return c.orig.SetDeadline(t)
}

func (c *ObfsWeChatUDPPacketConn) SetReadDeadline(t time.Time) error {
	return c.orig.SetReadDeadline(t)
}

func (c *ObfsWeChatUDPPacketConn) SetWriteDeadline(t time.Time) error {
	return c.orig.SetWriteDeadline(t)
}

func (c *ObfsWeChatUDPPacketConn) SetReadBuffer(bytes int) error {
	return c.orig.SetReadBuffer(bytes)
}

func (c *ObfsWeChatUDPPacketConn) SetWriteBuffer(bytes int) error {
	return c.orig.SetWriteBuffer(bytes)
}

func (c *ObfsWeChatUDPPacketConn) SyscallConn() (syscall.RawConn, error) {
	return c.orig.SyscallConn()
}

func (c *ObfsWeChatUDPPacketConn) File() (f *os.File, err error) {
	return c.orig.File()
}
