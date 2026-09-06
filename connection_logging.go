package quic

import (
	"net"

	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
	"github.com/quic-go/quic-go/qlog"
)

func toQlogFrame(frame wire.Frame) qlog.Frame { _ = "STUB: not implemented"; return *new(qlog.Frame) }

func toQlogAckFrame(f *wire.AckFrame) *qlog.AckFrame { _ = "STUB: not implemented"; return nil }

func (c *Conn) logLongHeaderPacket(p *longHeaderPacket, ecn protocol.ECN, datagramPayloadChecksum qlog.DatagramPayloadChecksum) {
	_ = "STUB: not implemented"
	return
}

func (c *Conn) logShortHeaderPacket(p shortHeaderPacket, ecn protocol.ECN, size protocol.ByteCount) {
	_ = "STUB: not implemented"
	return
}

func (c *Conn) logShortHeaderPacketWithDatagramPayloadChecksum(p shortHeaderPacket, ecn protocol.ECN, size protocol.ByteCount, isCoalesced bool, datagramPayloadChecksum qlog.DatagramPayloadChecksum) {
	_ = "STUB: not implemented"
	return
}

func (c *Conn) logCoalescedPacket(packet *coalescedPacket, ecn protocol.ECN) {
	_ = "STUB: not implemented"
	return
}

func (c *Conn) qlogTransportParameters(tp *wire.TransportParameters, sentBy protocol.Perspective, restore bool) {
	_ = "STUB: not implemented"
	return
}

func toQlogECN(ecn protocol.ECN) qlog.ECN {
	_ = "STUB: not implemented"
	//nolint:exhaustive // only need to handle the 3 valid values
	return *new(qlog.ECN)
}

func toQlogPacketType(pt protocol.PacketType) qlog.PacketType {
	_ = "STUB: not implemented"
	return *new(qlog.PacketType)
}

func toPathEndpointInfo(addr *net.UDPAddr) qlog.PathEndpointInfo {
	_ = "STUB: not implemented"
	return *new(qlog.PathEndpointInfo)
}

func startedConnectionEvent(local, remote *net.UDPAddr) qlog.StartedConnection {
	_ = "STUB: not implemented"
	return *new(qlog.StartedConnection)
}
