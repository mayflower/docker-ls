package connector

type semaphore chan int

func (s semaphore) Lock() {
	s <- 1
}

func (s semaphore) Unlock() {
	_ = <-s
}

func newSemaphore(limit uint) semaphore {
	if limit == 0 {
		limit = 1
	}

	return semaphore(make(chan int, limit))
}
