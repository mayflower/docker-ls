package main

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

func yamlToStdout(data interface{}) (err error) {
	marshalledData, err := yaml.Marshal(data)
	if err != nil {
		return
	}

	io.Copy(os.Stdout, bytes.NewReader(marshalledData))

	return
}
