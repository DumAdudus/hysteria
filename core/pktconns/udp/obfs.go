package udp

import (
	"net"

	"github.com/apernet/hysteria/core/pktconns/obfs"
	"github.com/valyala/bytebufferpool"
)

const (
	Protocol      = "udp"
	udpBufferSize = 1024 * 2
)

var (
	hintBuf                = make([]byte, udpBufferSize)
	_       net.PacketConn = &ObfsUDPPacketConn{}
)

type ObfsUDPPacketConn struct {
	*net.UDPConn
	obfs obfs.Obfuscator

	headerSize int
	header     []byte
}

func NewObfsUDPConn(orig *net.UDPConn, obfs obfs.Obfuscator) *ObfsUDPPacketConn {
	return &ObfsUDPPacketConn{
		UDPConn:    orig,
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
		n, addr, err := c.UDPConn.ReadFrom(poolBuf.Bytes())
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
		obfs, _ := c.obfs.(interface {
			ObfuscateOnBuffer([]byte, *bytebufferpool.ByteBuffer) int
		})
		obfs.ObfuscateOnBuffer(p, poolBuf)
	} else {
		poolBuf.Write(p)
	}

	n, err = c.UDPConn.WriteTo(poolBuf.Bytes(), addr)
	return
}

// This is a deliberate func to let ObfsUDPPacketConn not to be
// compatible with quic.OOBCapablePacketConn
func (c *ObfsUDPPacketConn) ReadMsgUDP() error {
	return nil
}
