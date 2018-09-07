package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func YamlToStdout(data interface{}) (err error) {
	marshalledData, err := yaml.Marshal(data)
	if err != nil {
		return
	}

	io.Copy(os.Stdout, bytes.NewReader(marshalledData))

	return
}

func JsonToStdout(data interface{}) (err error) {
	marshalledData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}

	io.Copy(os.Stdout, bytes.NewReader(marshalledData))
	fmt.Println()

	return
}

func SerializeToStdout(data interface{}, cfg *CliConfig) error {
	if cfg.JsonOutput {
		return JsonToStdout(data)
	} else {
		return YamlToStdout(data)
	}
}

func TemplateToStdout(data interface{}, name string) (err error) {
	var templateRepository TemplateRepository

	templateRepository, err = TemplateRepositoryFromConfig()
	if err != nil {
		return
	}

	template := templateRepository.Get(name)
	if template == nil {
		err = errors.New(fmt.Sprintf("no template with name '%s'", name))
		return
	}

	err = template.Execute(os.Stdout, data)
	return
}
