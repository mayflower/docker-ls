package main

import (
	"fmt"
	"os"
	"sync"
)

type progressIndicator struct {
	mutex sync.Mutex
	cfg   *Config
}

func (p *progressIndicator) Start(status string) {
	if p.cfg.progress {
		fmt.Fprint(os.Stderr, status+" ")
	}
}

func (p *progressIndicator) Progress() {
	if p.cfg.progress {
		p.mutex.Lock()
		fmt.Fprint(os.Stderr, ".")
		p.mutex.Unlock()
	}
}

func (p *progressIndicator) Finish(status string) {
	if p.cfg.progress {
		fmt.Fprintln(os.Stderr, " "+status)
	}
}

func NewProgressIndicator(cfg *Config) *progressIndicator {
	return &progressIndicator{
		cfg: cfg,
	}
}
