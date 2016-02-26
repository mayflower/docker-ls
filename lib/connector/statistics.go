package connector

import (
	"sync"
)

type Statistics interface {
	Requests() uint
	TokenCacheHitsAtApiLevel() uint
	TokenCacheMissesAtApiLevel() uint
	TokenCacheFailsAtApiLevel() uint
	TokenCacheHitsAtAuthLevel() uint
	TokenCacheMissesAtAuthLevel() uint
	TokenCacheFailsAtAuthLevel() uint
}

type statistics struct {
	requests               uint
	cacheMissesAtApiLevel  uint
	cacheHitsAtApiLevel    uint
	cacheFailsAtApiLevel   uint
	cacheHitsAtAuthLevel   uint
	cacheMissesAtAuthLevel uint
	cacheFailsAtAuthLevel  uint
	mutex                  sync.RWMutex
}

func (s *statistics) Requests() (r uint) {
	s.mutex.RLock()
	r = s.requests
	s.mutex.RUnlock()

	return
}

func (s *statistics) TokenCacheMissesAtApiLevel() (r uint) {
	s.mutex.RLock()
	r = s.cacheMissesAtApiLevel
	s.mutex.RUnlock()

	return
}

func (s *statistics) TokenCacheHitsAtApiLevel() (r uint) {
	s.mutex.RLock()
	r = s.cacheHitsAtApiLevel
	s.mutex.RUnlock()

	return
}

func (s *statistics) TokenCacheFailsAtApiLevel() (r uint) {
	s.mutex.RLock()
	r = s.cacheFailsAtApiLevel
	s.mutex.RUnlock()

	return
}

func (s *statistics) TokenCacheHitsAtAuthLevel() (r uint) {
	s.mutex.RLock()
	r = s.cacheHitsAtAuthLevel
	s.mutex.RUnlock()

	return
}

func (s *statistics) TokenCacheMissesAtAuthLevel() (r uint) {
	s.mutex.RLock()
	r = s.cacheMissesAtAuthLevel
	s.mutex.RUnlock()

	return
}

func (s *statistics) TokenCacheFailsAtAuthLevel() (r uint) {
	s.mutex.RLock()
	r = s.cacheFailsAtAuthLevel
	s.mutex.RUnlock()

	return
}

func (s *statistics) Request() {
	s.mutex.Lock()
	s.requests++
	s.mutex.Unlock()
}

func (s *statistics) CacheHitAtApiLevel() {
	s.mutex.Lock()
	s.cacheHitsAtApiLevel++
	s.mutex.Unlock()
}

func (s *statistics) CacheMissAtApiLevel() {
	s.mutex.Lock()
	s.cacheMissesAtApiLevel++
	s.mutex.Unlock()
}

func (s *statistics) CacheFailAtApiLevel() {
	s.mutex.Lock()
	s.cacheFailsAtApiLevel++
	s.mutex.Unlock()
}

func (s *statistics) CacheHitAtAuthLevel() {
	s.mutex.Lock()
	s.cacheHitsAtAuthLevel++
	s.mutex.Unlock()
}

func (s *statistics) CacheMissAtAuthLevel() {
	s.mutex.Lock()
	s.cacheMissesAtAuthLevel++
	s.mutex.Unlock()
}

func (s *statistics) CacheFailAtAuthLevel() {
	s.mutex.Lock()
	s.cacheFailsAtAuthLevel++
	s.mutex.Unlock()
}
