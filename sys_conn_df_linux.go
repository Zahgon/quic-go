//go:build linux

package quic

import (
	"syscall"
)

func setDF(rawConn syscall.RawConn) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func isSendMsgSizeErr(err error) bool { _ = "STUB: not implemented"; return false }

func isRecvMsgSizeErr(error) bool { _ = "STUB: not implemented"; return false }
