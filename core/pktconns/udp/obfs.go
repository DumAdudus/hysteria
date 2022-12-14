package udp

import (
	"net"
	"os"
	"syscall"
	"time"

	"github.com/apernet/hysteria/core/pktconns/obfs"
	"github.com/valyala/bytebufferpool"
)

const udpBufferSize = 2048

var (
	hintBuf                = make([]byte, udpBufferSize)
	_       net.PacketConn = &ObfsUDPPacketConn{}
)

type ObfsUDPPacketConn struct {
	orig *net.UDPConn
	obfs obfs.Obfuscator

	headerSize int
	header     []byte
}

func NewObfsUDPConn(orig *net.UDPConn, obfs obfs.Obfuscator) *ObfsUDPPacketConn {
	return &ObfsUDPPacketConn{
		orig:       orig,
		obfs:       obfs,
		headerSize: 0,
	}
}

func (c *ObfsUDPPacketConn) SetHeader(p []byte) {
	c.header = p
}

func (c *ObfsUDPPacketConn) SetHeaderSize(size int) {
	c.headerSize = size
}

func (c *ObfsUDPPacketConn) ReadFrom(p []byte) (int, net.Addr, error) {
	poolBuf := bytebufferpool.Get()
	defer bytebufferpool.Put(poolBuf)

	for {
		poolBuf.Set(hintBuf)
		n, addr, err := c.orig.ReadFrom(poolBuf.Bytes())
		if n <= c.headerSize {
			return 0, addr, err
		}
		payload := poolBuf.Bytes()[c.headerSize:n]
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

func (c *ObfsUDPPacketConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	poolBuf := bytebufferpool.Get()
	defer bytebufferpool.Put(poolBuf)
	poolBuf.Set(c.header)

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

func (c *ObfsUDPPacketConn) Close() error {
	return c.orig.Close()
}

func (c *ObfsUDPPacketConn) LocalAddr() net.Addr {
	return c.orig.LocalAddr()
}

func (c *ObfsUDPPacketConn) SetDeadline(t time.Time) error {
	return c.orig.SetDeadline(t)
}

func (c *ObfsUDPPacketConn) SetReadDeadline(t time.Time) error {
	return c.orig.SetReadDeadline(t)
}

func (c *ObfsUDPPacketConn) SetWriteDeadline(t time.Time) error {
	return c.orig.SetWriteDeadline(t)
}

func (c *ObfsUDPPacketConn) SetReadBuffer(bytes int) error {
	return c.orig.SetReadBuffer(bytes)
}

func (c *ObfsUDPPacketConn) SetWriteBuffer(bytes int) error {
	return c.orig.SetWriteBuffer(bytes)
}

func (c *ObfsUDPPacketConn) SyscallConn() (syscall.RawConn, error) {
	return c.orig.SyscallConn()
}

func (c *ObfsUDPPacketConn) File() (f *os.File, err error) {
	return c.orig.File()
}
