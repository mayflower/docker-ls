package util

import (
	"fmt"
	"os"
	"sync"
)

type progressIndicator struct {
	mutex sync.Mutex
	cfg   *CliConfig
}

func (p *progressIndicator) Start(status string) {
	if p.cfg.Progress {
		fmt.Fprint(os.Stderr, status+" ")
	}
}

func (p *progressIndicator) Progress() {
	if p.cfg.Progress {
		p.mutex.Lock()
		fmt.Fprint(os.Stderr, ".")
		p.mutex.Unlock()
	}
}

func (p *progressIndicator) Finish(status string) {
	if p.cfg.Progress {
		fmt.Fprintln(os.Stderr, " "+status)
	}
}

func NewProgressIndicator(cfg *CliConfig) *progressIndicator {
	return &progressIndicator{
		cfg: cfg,
	}
}
