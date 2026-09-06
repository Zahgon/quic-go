//go:build darwin || freebsd || linux || openbsd

package quic

import (
	"syscall"
)

func inspectReadBuffer(c syscall.RawConn) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func inspectWriteBuffer(c syscall.RawConn) (int, error) { _ = "STUB: not implemented"; return 0, nil }
