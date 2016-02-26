package connector

import (
	"sync"

	"github.com/mayflower/docker-ls/lib/auth"
)

type tokenCache struct {
	entries map[string]auth.Token
	mutex   sync.RWMutex
}

func (t *tokenCache) Get(hint string) (token auth.Token) {
	t.mutex.RLock()
	if value, cached := t.entries[hint]; cached {
		token = value
	}
	t.mutex.RUnlock()

	return
}

func (t *tokenCache) Set(hint string, token auth.Token) {
	t.mutex.Lock()
	t.entries[hint] = token
	t.mutex.Unlock()
}

func newTokenCache() *tokenCache {
	return &tokenCache{
		entries: make(map[string]auth.Token),
	}
}
