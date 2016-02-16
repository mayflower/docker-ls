package main

import (
	"flag"
)

type Config struct {
	recursionLevel uint
}

func (c *Config) bindToFlags(flags *flag.FlagSet) {
	flags.UintVar(&c.recursionLevel, "level", c.recursionLevel, "level of recursion")
}

func newConfig() *Config {
	return new(Config)
}
