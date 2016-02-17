package auth

import (
	"sort"
	"strings"
	"sync"
)

type challengeCacheKey struct {
	realm   string
	service string
	scopes  string
}

type tokenCache struct {
	entries map[challengeCacheKey]authResponse
	mutex   sync.RWMutex
}

func newCacheKey(challenge *Challenge) challengeCacheKey {
	scopes := append([]string(nil), challenge.scope...)
	sort.Sort(sort.StringSlice(scopes))

	return challengeCacheKey{
		realm:   challenge.realm.String(),
		service: challenge.service,
		scopes:  strings.Join(scopes, " "),
	}
}

func (c *tokenCache) Get(challenge *Challenge) (token string) {
	key := newCacheKey(challenge)

	c.mutex.RLock()
	if entry, cached := c.entries[key]; cached {
		token = entry.Token
	}
	c.mutex.RUnlock()

	return
}

func (c *tokenCache) Set(challenge *Challenge, response authResponse) {
	key := newCacheKey(challenge)

	c.mutex.Lock()
	c.entries[key] = response
	c.mutex.Unlock()
}

func newTokenCache() *tokenCache {
	return &tokenCache{
		entries: make(map[challengeCacheKey]authResponse),
	}
}
