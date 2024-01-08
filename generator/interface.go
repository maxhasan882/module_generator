package generator

import (
	"bytes"
	"path/filepath"
	"template/file"
	"text/template"
)

const interfaceTemplate = `
type {{ .Name }} interface {
{{ range .Methods }}{{ .Name }}({{ range $index, $param := .Params }}{{ if $index }}, {{ end }}{{ $param.Name }} {{ $param.Type }}{{ end }}) ({{ range $index, $return := .Returns }}{{ if $index }}, {{ end }}{{ $return }}{{ end }})
{{ end }}
}
`

// InterfaceConfig represents information about an interface in the YAML configuration.
type InterfaceConfig struct {
	Name    string   `yaml:"name"`
	Methods []Method `yaml:"methods"`
}

// Method represents information about a method in an interface.
type Method struct {
	Name    string   `yaml:"name"`
	Returns []string `yaml:"returns"`
	Params  []struct {
		Name string `yaml:"name"`
		Type string `yaml:"type"`
	} `yaml:"params"`
}

func generateInterfaceFileImports(data *Data) {
	templateData := struct {
		Module  string
		GotSson bool
	}{
		Module: data.Module,
	}

	for _, item := range data.Interfaces {
		for _, rt := range item.Methods {
			for _, param := range rt.Params {
				if param.Type == "sson.D" || param.Type == "sson.M" || param.Type == "sson.E" {
					templateData.GotSson = true
				}
			}
		}
	}

	var buf bytes.Buffer
	tmpl, err := template.New("interfaceHeader").Parse(interfaceHeaderTemplate)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&buf, templateData)
	if err != nil {
		panic(err)
	}

	formattedCode, err := formatGoCode(buf.String())
	if err != nil {
		panic(err)
	}

	err = file.CreateOrUpdateModule(filepath.Join([]string{"domain", "repository"}...), data.FileName, formattedCode)
	if err != nil {
		panic(err)
	}
}

func generateInterface(data *Data) {
	generateInterfaceFileImports(data)
	configs := data.Interfaces
	for _, config := range configs {
		tmpl, err := template.New("interface").Parse(interfaceTemplate)
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, config)
		if err != nil {
			panic(err)
		}
		formattedCode, err := formatGoCode(buf.String())
		if err != nil {
			panic(err)
		}
		err = file.CreateOrUpdateModule(filepath.Join([]string{"domain", "repository"}...), data.FileName, formattedCode)
	}
}
