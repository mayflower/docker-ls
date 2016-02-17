package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func jsonToStdout(data interface{}) (err error) {
	marshalledData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}

	io.Copy(os.Stdout, bytes.NewReader(marshalledData))
	fmt.Println()

	return
}

func serializeToStdout(data interface{}, cfg *Config) error {
	if cfg.jsonOutput {
		return jsonToStdout(data)
	} else {
		return yamlToStdout(data)
	}
}
