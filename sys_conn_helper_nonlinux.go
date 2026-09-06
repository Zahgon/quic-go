//go:build !linux

package quic

func forceSetReceiveBuffer(c any, bytes int) error { _ = "STUB: not implemented"; return nil }
func forceSetSendBuffer(c any, bytes int) error    { _ = "STUB: not implemented"; return nil }

func appendUDPSegmentSizeMsg([]byte, uint16) []byte { _ = "STUB: not implemented"; return nil }
func isGSOError(error) bool                         { _ = "STUB: not implemented"; return false }
func isPermissionError(err error) bool              { _ = "STUB: not implemented"; return false }
