package main

type command interface {
	execute(argv []string) error
}
