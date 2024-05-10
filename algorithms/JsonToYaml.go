package main

import (
	"github.com/ghodss/yaml"
)

func convert(json string) string {

	b, _ := yaml.JSONToYAML([]byte(json))
	return string(b)
}
