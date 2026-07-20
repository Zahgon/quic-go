//go:build windows

package quic

import (
	"syscall"
)

const (

	//nolint:stylecheck
	IP_DONTFRAGMENT = 14

	//nolint:stylecheck
	IPV6_DONTFRAG = 14
)

func setDF(rawConn syscall.RawConn) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func isSendMsgSizeErr(err error) bool { _ = "STUB: not implemented"; return false }

func isRecvMsgSizeErr(err error) bool { _ = "STUB: not implemented"; return false }
