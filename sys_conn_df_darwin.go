//go:build darwin

package quic

import (
	"syscall"
)

const (
	macOSVersion11 = 20
	macOSVersion15 = 24
)

func setDF(rawConn syscall.RawConn) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func isSendMsgSizeErr(err error) bool { _ = "STUB: not implemented"; return false }

func isRecvMsgSizeErr(error) bool { _ = "STUB: not implemented"; return false }

func getMacOSVersion() (int, error) { _ = "STUB: not implemented"; return 0, nil }
