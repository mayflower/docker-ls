package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.SetOutput(ioutil.Discard)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
