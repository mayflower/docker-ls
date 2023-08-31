package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/template"

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
	if cfg.TemplateSource != "" {
		return templateSourceToStdout(data, cfg)
	} else if cfg.Template != "" {
		return namedTemplateToStdout(data, cfg)
	} else if cfg.JsonOutput {
		return JsonToStdout(data)
	} else {
		return YamlToStdout(data)
	}
}

func namedTemplateToStdout(data interface{}, cfg *CliConfig) (err error) {
	template := cfg.templateRepository.Get(cfg.Template)
	if template == nil {
		err = fmt.Errorf("no template with name '%s'", cfg.Template)
		return
	}

	err = template.Execute(os.Stdout, data)
	return
}

func templateSourceToStdout(data interface{}, cfg *CliConfig) (err error) {
	template := template.New("inline")

	_, err = template.Parse(cfg.TemplateSource)
	if err != nil {
		return
	}

	err = template.Execute(os.Stdout, data)
	return
}
