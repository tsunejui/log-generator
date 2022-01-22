package template

import (
	"bytes"
	"fmt"
	ltmp "log-generator/template"
	"text/template"
)

type Log struct{}

type LogConfig struct {
	Name   string
	Source string
	Data   interface{}
}

func New() *Log {
	return &Log{}
}

const nameJsonTemplate = "Json Template"

func (l *Log) GetJsonFormatTemplate(timez, level string) (string, error) {
	tmp, err := genTemplate(&LogConfig{
		Name:   nameJsonTemplate,
		Source: ltmp.JsonTemplate(),
		Data: map[string]interface{}{
			"Time":  timez,
			"Level": level,
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate a template: %v", err)
	}
	return tmp, nil
}

const nameTextTemplate = "Text Template"

func (l *Log) GetTextFormatTemplate(timez, level string) (string, error) {
	tmp, err := genTemplate(&LogConfig{
		Name:   nameTextTemplate,
		Source: ltmp.TextTemplate(),
		Data: map[string]interface{}{
			"Time":  timez,
			"Level": level,
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate a template: %v", err)
	}
	return tmp, nil
}

func genTemplate(config *LogConfig) (string, error) {
	t := template.New(config.Name)
	tmp, err := t.Parse(config.Source)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	var tpl bytes.Buffer
	if err := tmp.Execute(&tpl, config.Data); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return tpl.String(), nil
}
