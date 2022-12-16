//go:build debug
// +build debug

package udp

import (
	"fmt"
	"net"
	"net/netip"
	"os"
	"runtime"
	"syscall"
	"time"

	"github.com/apernet/hysteria/core/pktconns/obfs"
	"github.com/valyala/bytebufferpool"
)

func (c *ObfsUDPPacketConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *net.UDPAddr, err error) {
	poolBuf := bytebufferpool.Get()
	defer bytebufferpool.Put(poolBuf)
	var ap netip.AddrPort

	for {
		poolBuf.Set(hintBuf)

		n, oobn, flags, ap, err = c.UDPConn.ReadMsgUDPAddrPort(poolBuf.Bytes(), oob)
		if ap.IsValid() {
			addr = net.UDPAddrFromAddrPort(ap)
		}

		if n <= c.headerSize {
			n = 0
			return
		}
		payload := poolBuf.Bytes()[c.headerSize:n]
		if c.obfs != nil {
			n = c.obfs.Deobfuscate(payload, b)
		} else {
			n = copy(b, payload)
		}
		if n > 0 || err != nil {
			// Valid packet
			return
		}
	}

	return
}

func (c *ObfsUDPPacketConn) WriteMsgUDP(b, oob []byte, addr *net.UDPAddr) (n, oobn int, err error) {
	poolBuf := bytebufferpool.Get()
	defer bytebufferpool.Put(poolBuf)
	poolBuf.Set(c.header)

	if c.obfs != nil {
		xplusObfs, _ := c.obfs.(*obfs.XPlusObfuscator) // Currently there's only this XPlus obfs
		xplusObfs.ObfuscateOnBuffer(b, poolBuf)
	} else {
		poolBuf.Write(b)
	}

	n, oobn, err = c.UDPConn.WriteMsgUDP(b, oob, addr)
	return
}

func (c *ObfsUDPPacketConn) Read(b []byte) (int, error) {
	printFuncName()
	return c.UDPConn.Read(b)
}

func (c *ObfsUDPPacketConn) Write(b []byte) (int, error) {
	printFuncName()
	return c.UDPConn.Write(b)
}

func (c *ObfsUDPPacketConn) ReadFromUDP(b []byte) (int, *net.UDPAddr, error) {
	printFuncName()
	return c.UDPConn.ReadFromUDP(b)
}

func (c *ObfsUDPPacketConn) ReadFromUDPAddrPort(b []byte) (int, netip.AddrPort, error) {
	printFuncName()
	return c.UDPConn.ReadFromUDPAddrPort(b)
}

func (c *ObfsUDPPacketConn) ReadMsgUDPAddrPort(b, oob []byte) (int, int, int, netip.AddrPort, error) {
	printFuncName()
	return c.UDPConn.ReadMsgUDPAddrPort(b, oob)
}

func (c *ObfsUDPPacketConn) WriteToUDP(b []byte, addr *net.UDPAddr) (int, error) {
	printFuncName()
	return c.UDPConn.WriteToUDP(b, addr)
}

func (c *ObfsUDPPacketConn) WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error) {
	printFuncName()
	return c.UDPConn.WriteToUDPAddrPort(b, addr)
}

func (c *ObfsUDPPacketConn) WriteMsgUDPAddrPort(b, oob []byte, addr netip.AddrPort) (int, int, error) {
	printFuncName()
	return c.UDPConn.WriteMsgUDPAddrPort(b, oob, addr)
}

func (c *ObfsUDPPacketConn) Close() error {
	printFuncName()
	return c.UDPConn.Close()
}

func (c *ObfsUDPPacketConn) LocalAddr() net.Addr {
	printFuncName()
	return c.UDPConn.LocalAddr()
}

func (c *ObfsUDPPacketConn) RemoteAddr() net.Addr {
	printFuncName()
	return c.UDPConn.RemoteAddr()
}

func (c *ObfsUDPPacketConn) SetDeadline(t time.Time) error {
	printFuncName()
	return c.UDPConn.SetDeadline(t)
}

func (c *ObfsUDPPacketConn) SetReadDeadline(t time.Time) error {
	printFuncName()
	return c.UDPConn.SetReadDeadline(t)
}

func (c *ObfsUDPPacketConn) SetWriteDeadline(t time.Time) error {
	printFuncName()
	return c.UDPConn.SetWriteDeadline(t)
}

func (c *ObfsUDPPacketConn) SetReadBuffer(bytes int) error {
	printFuncName()
	return c.UDPConn.SetReadBuffer(bytes)
}

func (c *ObfsUDPPacketConn) SetWriteBuffer(bytes int) error {
	printFuncName()
	return c.UDPConn.SetWriteBuffer(bytes)
}

func (c *ObfsUDPPacketConn) SyscallConn() (syscall.RawConn, error) {
	printFuncName()
	return c.UDPConn.SyscallConn()
}

func (c *ObfsUDPPacketConn) File() (f *os.File, err error) {
	printFuncName()
	return c.UDPConn.File()
}

func printFuncName() {
	pc, _, _, _ := runtime.Caller(1)
	fmt.Printf("%s\n", runtime.FuncForPC(pc).Name())
}
