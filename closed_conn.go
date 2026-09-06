package quic

import (
	"net"
	"sync/atomic"

	"github.com/quic-go/quic-go/internal/utils"
)

type closedLocalConn struct {
	counter atomic.Uint32
	logger  utils.Logger

	sendPacket func(net.Addr, packetInfo)
}

var _ packetHandler = &closedLocalConn{}

func newClosedLocalConn(sendPacket func(net.Addr, packetInfo), logger utils.Logger) packetHandler {
	_ = "STUB: not implemented"
	return *new(packetHandler)
}

func (c *closedLocalConn) handlePacket(p receivedPacket) { _ = "STUB: not implemented"; return }

func (c *closedLocalConn) destroy(error) { _ = "STUB: not implemented"; return }
func (c *closedLocalConn) closeWithTransportError(TransportErrorCode) {
	_ = "STUB: not implemented"
	return
}

type closedRemoteConn struct{}

var _ packetHandler = &closedRemoteConn{}

func newClosedRemoteConn() packetHandler { _ = "STUB: not implemented"; return *new(packetHandler) }

func (c *closedRemoteConn) handlePacket(receivedPacket) { _ = "STUB: not implemented"; return }
func (c *closedRemoteConn) destroy(error)               { _ = "STUB: not implemented"; return }
func (c *closedRemoteConn) closeWithTransportError(TransportErrorCode) {
	_ = "STUB: not implemented"
	return
}
