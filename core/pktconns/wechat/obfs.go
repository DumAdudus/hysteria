package wechat

import (
	"encoding/binary"
	"math/rand"
	"net"

	"github.com/apernet/hysteria/core/pktconns/obfs"
	"github.com/apernet/hysteria/core/pktconns/udp"
)

const (
	headerSize = 13
)

var initHeader = [headerSize]byte{0xa1, 0x08, 0xff, 0xff, 0xff, 0xff, 0x00, 0x10, 0x11, 0x18, 0x30, 0x22, 0x30}

// ObfsWeChatUDPPacketConn is still a UDP packet conn, but it adds WeChat video call header to each packet.
// Obfs in this case can be nil
type ObfsWeChatUDPPacketConn struct {
	*udp.ObfsUDPPacketConn
	sn uint32
}

func NewObfsWeChatUDPConn(orig *net.UDPConn, obfs obfs.Obfuscator) *ObfsWeChatUDPPacketConn {
	conn := &ObfsWeChatUDPPacketConn{
		ObfsUDPPacketConn: udp.NewObfsUDPConn(orig, obfs),
		sn:                rand.Uint32() & 0xFFF,
	}
	conn.ObfsUDPPacketConn.SetHeaderSize(headerSize)
	return conn
}

func (c *ObfsWeChatUDPPacketConn) ReadFrom(p []byte) (int, net.Addr, error) {
	return c.ObfsUDPPacketConn.ReadFrom(p)
}

func (c *ObfsWeChatUDPPacketConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	header := initHeader[:]
	// Set SN, we don't care what is it, so no lock
	sn := header[2:6]
	binary.BigEndian.PutUint32(sn, c.sn)
	c.sn++

	c.ObfsUDPPacketConn.SetHeader(header)

	return c.ObfsUDPPacketConn.WriteTo(p, addr)
}
