package quic

import (
	"sync"

	list "github.com/quic-go/quic-go/internal/utils/linkedlist"
)

type singleOriginTokenStore struct {
	tokens []*ClientToken
	len    int
	p      int
}

func newSingleOriginTokenStore(size int) *singleOriginTokenStore {
	_ = "STUB: not implemented"
	return nil
}

func (s *singleOriginTokenStore) Add(token *ClientToken) { _ = "STUB: not implemented"; return }

func (s *singleOriginTokenStore) Pop() *ClientToken { _ = "STUB: not implemented"; return nil }

func (s *singleOriginTokenStore) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *singleOriginTokenStore) index(i int) int { _ = "STUB: not implemented"; return 0 }

type lruTokenStoreEntry struct {
	key   string
	cache *singleOriginTokenStore
}

type lruTokenStore struct {
	mutex sync.Mutex

	m                map[string]*list.Element[*lruTokenStoreEntry]
	q                *list.List[*lruTokenStoreEntry]
	capacity         int
	singleOriginSize int
}

var _ TokenStore = &lruTokenStore{}

func NewLRUTokenStore(maxOrigins, tokensPerOrigin int) TokenStore {
	_ = "STUB: not implemented"
	return *new(TokenStore)
}

func (s *lruTokenStore) Put(key string, token *ClientToken) { _ = "STUB: not implemented"; return }

func (s *lruTokenStore) Pop(key string) *ClientToken { _ = "STUB: not implemented"; return nil }
