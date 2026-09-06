//go:build !linux && !windows && !darwin

package quic

import (
	"syscall"
)

func setDF(syscall.RawConn) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func isSendMsgSizeErr(err error) bool { _ = "STUB: not implemented"; return false }

func isRecvMsgSizeErr(err error) bool { _ = "STUB: not implemented"; return false }
