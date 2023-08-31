package util

import (
	"fmt"
	"text/template"

	"github.com/spf13/viper"
)

type TemplateRepository interface {
	Get(string) *template.Template
	Has(string) bool
}

type templateRepository map[string]*template.Template

func (t templateRepository) Get(name string) *template.Template {
	return t[name]
}

func (t templateRepository) Has(name string) bool {
	_, ok := t[name]

	return ok
}

func TemplateRepositoryFromConfig() (t TemplateRepository, err error) {
	var compiled templateRepository = make(map[string]*template.Template)
	t = compiled

	if !viper.IsSet("templates") {
		return
	}

	source := make(map[string]string)

	err = viper.UnmarshalKey("templates", &source)
	if err != nil {
		return
	}

	for name, s := range source {
		tpl := template.New(name)

		if _, e := tpl.Parse(s); e != nil {
			err = fmt.Errorf("syntax error in template '%s': %v", name, e)
			return
		}

		compiled[name] = tpl
	}

	return
}
