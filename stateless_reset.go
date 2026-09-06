package quic

import (
	"hash"
	"sync"

	"github.com/quic-go/quic-go/internal/protocol"
)

type statelessResetter struct {
	mx sync.Mutex
	h  hash.Hash
}

func newStatelessResetter(key *StatelessResetKey) *statelessResetter {
	_ = "STUB: not implemented"
	return nil
}

func (r *statelessResetter) GetStatelessResetToken(connID protocol.ConnectionID) protocol.StatelessResetToken {
	_ = "STUB: not implemented"
	return *new(protocol.StatelessResetToken)
}
