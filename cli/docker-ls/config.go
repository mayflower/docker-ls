package main

import (
	"flag"
)

type Config struct {
	recursionLevel uint
	statistics     bool
}

func (c *Config) bindToFlags(flags *flag.FlagSet) {
	flags.UintVar(&c.recursionLevel, "level", c.recursionLevel, "level of recursion")
	flags.BoolVar(&c.statistics, "statistics", false, "show connection statistics")
}

func newConfig() *Config {
	return new(Config)
}
