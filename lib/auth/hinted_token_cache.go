package auth

import (
	"sync"
)

type hintedTokenCache struct {
	entries map[string]authResponse
	mutex   sync.RWMutex
}

func (c *hintedTokenCache) Get(hint string) (token string) {
	c.mutex.RLock()
	if entry, cached := c.entries[hint]; cached {
		token = entry.Token
	}
	c.mutex.RUnlock()

	return
}

func (c *hintedTokenCache) Set(hint string, response authResponse) {
	c.mutex.Lock()
	c.entries[hint] = response
	c.mutex.Unlock()
}

func newHintedTokenCache() *hintedTokenCache {
	return &hintedTokenCache{
		entries: make(map[string]authResponse),
	}
}
